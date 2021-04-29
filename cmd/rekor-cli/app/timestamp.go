//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"encoding/asn1"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"

	"github.com/go-openapi/strfmt"
	"github.com/sassoftware/relic/lib/pkcs9"
	"github.com/sigstore/rekor/cmd/rekor-cli/app/format"
	"github.com/sigstore/rekor/pkg/generated/client/timestamp"
	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/sigstore/rekor/pkg/log"
	"github.com/sigstore/rekor/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type fileFlag struct {
	value string
}

func (f *fileFlag) String() string {
	return f.value
}

func (f *fileFlag) Set(s string) error {
	if s == "" {
		return errors.New("flag must be specified")
	}
	if _, err := os.Stat(filepath.Clean(s)); os.IsNotExist(err) {
		return err
	}
	f.value = s
	return nil
}

func (f *fileFlag) Type() string {
	return "fileFlag"
}

func addTimestampFlags(cmd *cobra.Command) error {
	cmd.Flags().Var(&fileFlag{}, "request", "path to an RFC 3161 timestamp request")

	cmd.Flags().Var(&fileFlag{}, "file", "path to a file containing the message to timestamp")

	cmd.Flags().Var(&urlFlag{}, "url", "URL to a timestamp authority, e.g. https://freetsa.org/tsr")

	// TODO: Add a flag to upload the timestamp response to the log and return the log index.
	// TODO: Add data to validate response.
	return nil
}

func validateTimestampFlags() error {
	requestStr := viper.GetString("request")
	fileStr := viper.GetString("file")
	urlStr := viper.GetString("url")
	if requestStr == "" && fileStr == "" {
		return errors.New("request or file must be specified")
	}
	// TODO: Make this optional when rekor can produce RFC 3161 os JSON signatures.
	if urlStr == "" {
		return errors.New("URL must be specified")
	}
	return nil
}

func createRequestFromFlags() (*pkcs9.TimeStampReq, error) {
	requestStr := viper.GetString("request")
	var timestampReq *pkcs9.TimeStampReq
	if requestStr == "" {
		fileStr := viper.GetString("file")
		fileBytes, err := ioutil.ReadFile(filepath.Clean(fileStr))
		if err != nil {
			return nil, fmt.Errorf("error reading request from file: %w", err)
		}
		timestampReq, err = util.TimestampRequestFromData(fileBytes)
		if err != nil {
			return nil, fmt.Errorf("error creating timestamp request: %w", err)
		}
	} else {
		rawBytes, err := ioutil.ReadFile(filepath.Clean(requestStr))
		if err != nil {
			return nil, fmt.Errorf("error reading request from file: %w", err)
		}
		timestampReq, err = util.ParseTimestampRequest(rawBytes)
		if err != nil {
			return nil, fmt.Errorf("error parsing timestamp request: %w", err)
		}
	}
	return timestampReq, nil
}

type timestampCmdOutput struct {
	Location string
}

func (t *timestampCmdOutput) String() string {
	return fmt.Sprintf(`
Wrote response to: %v
`, t.Location)
}

var timestampCmd = &cobra.Command{
	Use:   "timestamp",
	Short: "Rekor timestamp command",
	Long:  "Retrieves an RFC 3161 timestamp response from a TSA with a specified URL. Note: the user must verify the RFC 3161 response with the TSA's certificates",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			log.Logger.Fatal("Error initializing cmd line args: ", err)
		}
		if err := validateTimestampFlags(); err != nil {
			log.Logger.Error(err)
			_ = cmd.Help()
			return err
		}
		return nil
	},
	Run: format.WrapCmd(func(args []string) (interface{}, error) {
		rekorClient, err := GetRekorClient(viper.GetString("rekor_server"))
		if err != nil {
			return nil, err
		}

		timestampReq, err := createRequestFromFlags()
		if err != nil {
			return nil, err
		}
		requestBytes, err := asn1.Marshal(*timestampReq)
		if err != nil {
			return nil, err
		}

		urlStr := viper.GetString("url")
		tsaURL, err := url.Parse(urlStr)
		if err != nil || !tsaURL.IsAbs() {
			return nil, fmt.Errorf("error parsing URL %w", err)
		}

		params := timestamp.NewGetTimestampResponseParams()
		params.Query = &models.TimestampRequest{}
		params.Query.URL = strfmt.URI(urlStr)
		base64Request := make([]byte, base64.StdEncoding.EncodedLen(len(requestBytes)))
		base64.StdEncoding.Encode(base64Request, requestBytes)
		params.Query.RfcRequest = strfmt.Base64(requestBytes)

		resp, err := rekorClient.Timestamp.GetTimestampResponse(params)
		if err != nil {
			return nil, err
		}
		body, err := base64.StdEncoding.DecodeString(resp.Payload.RfcResponse.String())
		if err != nil {
			return nil, err
		}

		// Sanity check response and check if the TimeStampToken was successfully created
		_, err = timestampReq.ParseResponse(body)
		if err != nil {
			return nil, err
		}

		// Write response to file
		f, err := os.Create("response.tsr")
		if err != nil {
			return nil, err
		}
		defer f.Close()

		if _, err := f.Write(body); err != nil {
			return nil, err
		}
		if err := f.Sync(); err != nil {
			return nil, err
		}
		return &timestampCmdOutput{
			Location: f.Name(),
		}, nil
	}),
}

func init() {
	if err := addTimestampFlags(timestampCmd); err != nil {
		log.Logger.Fatal("Error parsing cmd line args: ", err)
	}

	rootCmd.AddCommand(timestampCmd)
}
