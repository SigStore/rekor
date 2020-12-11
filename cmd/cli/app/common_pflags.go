/*
Copyright © 2020 Bob Callaway <bcallawa@redhat.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/projectrekor/rekor/pkg/generated/models"
	rekord_v001 "github.com/projectrekor/rekor/pkg/types/rekord/v0.0.1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addArtifactPFlags(cmd *cobra.Command) error {
	cmd.Flags().Var(&fileOrURLFlag{}, "signature", "path or URL to detached signature file")

	cmd.Flags().Var(&fileOrURLFlag{}, "public-key", "path or URL to public key file")

	cmd.Flags().Var(&fileOrURLFlag{}, "artifact", "path or URL to artifact file")

	cmd.Flags().Var(&fileOrURLFlag{}, "rekord", "path or URL to Rekor rekord file")

	cmd.Flags().Var(&shaFlag{}, "sha", "the sha of the artifact")
	return nil
}

func validateArtifactPFlags() error {
	// we will need artifact, public-key, signature, and potentially SHA
	rekord := viper.GetString("rekord")

	artifact := fileOrURLFlag{}
	artifactStr := viper.GetString("artifact")
	if artifactStr != "" {
		if err := artifact.Set(artifactStr); err != nil {
			return err
		}
	}

	signature := viper.GetString("signature")
	publicKey := viper.GetString("public-key")
	sha := viper.GetString("sha")

	if rekord == "" && artifact.String() == "" {
		return errors.New("either 'rekord' or 'artifact' must be specified")
	}

	if rekord == "" {
		if artifact.IsURL && sha == "" {
			return errors.New("a valid SHA hash must be specified when specifying a URL for --artifact")
		}
		if signature == "" {
			return errors.New("--signature is required when --artifact is used")
		}
		if publicKey == "" {
			return errors.New("--public-key is required when --artifact is used")
		}
	}

	return nil
}

func CreateRekordFromPFlags() (models.ProposedEntry, error) {
	//TODO: how to select version of item to create
	returnVal := models.Rekord{}
	re := new(rekord_v001.V001Entry)

	rekord := viper.GetString("rekord")
	if rekord != "" {
		var rekordBytes []byte
		rekordURL, err := url.Parse(rekord)
		if err == nil && rekordURL.IsAbs() {
			rekordResp, err := http.Get(rekord)
			if err != nil {
				return nil, fmt.Errorf("error fetching 'rekord': %w", err)
			}
			defer rekordResp.Body.Close()
			rekordBytes, err = ioutil.ReadAll(rekordResp.Body)
			if err != nil {
				return nil, fmt.Errorf("error fetching 'rekord': %w", err)
			}
		} else {
			rekordBytes, err = ioutil.ReadFile(filepath.Clean(rekord))
			if err != nil {
				return nil, fmt.Errorf("error processing 'rekord' file: %w", err)
			}
		}
		if err := json.Unmarshal(rekordBytes, &returnVal); err != nil {
			return nil, fmt.Errorf("error parsing rekord file: %w", err)
		}
	} else {
		// we will need artifact, public-key, signature, and potentially SHA
		re.RekordObj.Data = &models.RekordV001SchemaData{}

		artifact := viper.GetString("artifact")
		dataURL, err := url.Parse(artifact)
		if err == nil && dataURL.IsAbs() {
			re.RekordObj.Data.URL = strfmt.URI(artifact)
			re.RekordObj.Data.Hash = &models.RekordV001SchemaDataHash{}
			re.RekordObj.Data.Hash.Algorithm = swag.String("sha256")
			re.RekordObj.Data.Hash.Value = swag.String(viper.GetString("sha"))
		} else {
			artifactBytes, err := ioutil.ReadFile(filepath.Clean(artifact))
			if err != nil {
				return nil, fmt.Errorf("error reading artifact file: %w", err)
			}
			re.RekordObj.Data.Content = strfmt.Base64(artifactBytes)
		}

		re.RekordObj.Signature = &models.RekordV001SchemaSignature{}
		re.RekordObj.Signature.Format = models.RekordV001SchemaSignatureFormatPgp
		signature := viper.GetString("signature")
		sigURL, err := url.Parse(signature)
		if err == nil && sigURL.IsAbs() {
			re.RekordObj.Signature.URL = strfmt.URI(signature)
		} else {
			signatureBytes, err := ioutil.ReadFile(filepath.Clean(signature))
			if err != nil {
				return nil, fmt.Errorf("error reading signature file: %w", err)
			}
			re.RekordObj.Signature.Content = strfmt.Base64(signatureBytes)
		}

		re.RekordObj.Signature.PublicKey = &models.RekordV001SchemaSignaturePublicKey{}
		publicKey := viper.GetString("public-key")
		keyURL, err := url.Parse(publicKey)
		if err == nil && keyURL.IsAbs() {
			re.RekordObj.Signature.PublicKey.URL = strfmt.URI(publicKey)
		} else {
			keyBytes, err := ioutil.ReadFile(filepath.Clean(publicKey))
			if err != nil {
				return nil, fmt.Errorf("error reading public key file: %w", err)
			}
			re.RekordObj.Signature.PublicKey.Content = strfmt.Base64(keyBytes)
		}

		if re.HasExternalEntities() {
			if err := re.FetchExternalEntities(context.Background()); err != nil {
				return nil, fmt.Errorf("error retrieving external entities: %v", err)
			}
		}

		returnVal.APIVersion = swag.String(re.APIVersion())
		returnVal.Spec = re.RekordObj
	}

	if err := returnVal.Validate(nil); err != nil {
		return nil, err
	}

	return &returnVal, nil
}

type fileOrURLFlag struct {
	value string
	IsURL bool
}

func (f *fileOrURLFlag) String() string {
	return f.value
}

func (f *fileOrURLFlag) Set(s string) error {
	if s == "" {
		return errors.New("flag must be specified")
	}
	if _, err := os.Stat(filepath.Clean(s)); os.IsNotExist(err) {
		u := urlFlag{}
		if err := u.Set(s); err != nil {
			return err
		}
		f.IsURL = true
	}
	f.value = s
	return nil
}

func (f *fileOrURLFlag) Type() string {
	return "fileOrURLFlag"
}

type shaFlag struct {
	hash string
}

func (s *shaFlag) String() string {
	return s.hash
}

func (s *shaFlag) Set(v string) error {
	if v == "" {
		return errors.New("flag must be specified")
	}
	if _, err := hex.DecodeString(v); (err != nil) || (len(v) != 64) {
		if err == nil {
			err = errors.New("invalid length for value")
		}
		return fmt.Errorf("value specified is invalid: %w", err)
	}
	s.hash = v
	return nil
}

func (s *shaFlag) Type() string {
	return "sha"
}

func addUUIDPFlags(cmd *cobra.Command, required bool) error {
	cmd.Flags().Var(&shaFlag{}, "uuid", "UUID of entry in transparency log (if known)")
	if required {
		if err := cmd.MarkFlagRequired("uuid"); err != nil {
			return err
		}
	}
	return nil
}
