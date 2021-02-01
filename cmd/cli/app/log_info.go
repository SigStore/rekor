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
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/projectrekor/rekor/cmd/cli/app/state"

	"github.com/google/trillian"
	tclient "github.com/google/trillian/client"
	tcrypto "github.com/google/trillian/crypto"
	"github.com/google/trillian/merkle"
	"github.com/google/trillian/merkle/rfc6962"

	"github.com/projectrekor/rekor/cmd/cli/app/format"
	"github.com/projectrekor/rekor/pkg/generated/client/tlog"
	"github.com/projectrekor/rekor/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type logInfoCmdOutput struct {
	TreeSize int64
	RootHash string
}

func (l *logInfoCmdOutput) String() string {
	// Verification is always successful if we return an object.
	return fmt.Sprintf("Verification Successful!\nTree Size: %v\nRoot Hash: %s\n", l.TreeSize, l.RootHash)
}

// logInfoCmd represents the current information about the transparency log
var logInfoCmd = &cobra.Command{
	Use:   "loginfo",
	Short: "Rekor loginfo command",
	Long:  `Prints info about the transparency log`,
	Run: format.WrapCmd(func(args []string) (interface{}, error) {
		serverURL := viper.GetString("rekor_server")
		rekorClient, err := GetRekorClient(serverURL)
		if err != nil {
			return nil, err
		}

		result, err := rekorClient.Tlog.GetLogInfo(nil)
		if err != nil {
			return nil, err
		}

		logInfo := result.GetPayload()
		cmdOutput := &logInfoCmdOutput{
			TreeSize: *logInfo.TreeSize,
			RootHash: *logInfo.RootHash,
		}

		keyHint, err := base64.StdEncoding.DecodeString(logInfo.SignedTreeHead.KeyHint.String())
		if err != nil {
			return nil, err
		}
		logRoot, err := base64.StdEncoding.DecodeString(logInfo.SignedTreeHead.LogRoot.String())
		if err != nil {
			return nil, err
		}
		signature, err := base64.StdEncoding.DecodeString(logInfo.SignedTreeHead.Signature.String())
		if err != nil {
			return nil, err
		}
		sth := trillian.SignedLogRoot{
			KeyHint:          keyHint,
			LogRoot:          logRoot,
			LogRootSignature: signature,
		}

		publicKey := viper.GetString("rekor_server_public_key")
		if publicKey == "" {
			// fetch key from server
			keyResp, err := rekorClient.Tlog.GetPublicKey(nil)
			if err != nil {
				return nil, err
			}
			publicKey = keyResp.Payload
		}

		block, _ := pem.Decode([]byte(publicKey))
		if block == nil {
			return nil, errors.New("failed to decode public key of server")
		}

		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		verifier := tclient.NewLogVerifier(rfc6962.DefaultHasher, pub, crypto.SHA256)
		lr, err := tcrypto.VerifySignedLogRoot(verifier.PubKey, verifier.SigHash, &sth)
		if err != nil {
			return nil, err
		}

		oldState := state.Load(serverURL)
		if oldState != nil {
			log.CliLogger.Infof("Found previous log state, proving consistency between %d and %d", oldState.TreeSize, lr.TreeSize)
			params := tlog.NewGetLogProofParams()
			firstSize := int64(oldState.TreeSize)
			params.FirstSize = &firstSize
			params.LastSize = int64(lr.TreeSize)
			proof, err := rekorClient.Tlog.GetLogProof(params)
			if err != nil {
				return nil, err
			}
			hashes := [][]byte{}
			for _, h := range proof.Payload.Hashes {
				b, _ := hex.DecodeString(h)
				hashes = append(hashes, b)
			}
			v := merkle.NewLogVerifier(rfc6962.DefaultHasher)
			if err := v.VerifyConsistencyProof(firstSize, int64(lr.TreeSize), oldState.RootHash,
				lr.RootHash, hashes); err != nil {
				return nil, err
			}
			log.CliLogger.Infof("Consistency proof valid!")
		} else {
			log.CliLogger.Infof("No previous log state stored, unable to prove consistency")
		}

		if viper.GetBool("store_tree_state") {
			if err := state.Dump(serverURL, lr); err != nil {
				log.CliLogger.Infof("Unable to store previous state: %v", err)
			}
		}
		return cmdOutput, nil
	}),
}

func init() {
	rootCmd.AddCommand(logInfoCmd)
}
