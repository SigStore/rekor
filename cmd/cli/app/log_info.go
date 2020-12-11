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
	"fmt"

	"github.com/projectrekor/rekor/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// logInfoCmd represents the current information about the transparency log
var logInfoCmd = &cobra.Command{
	Use:   "loginfo",
	Short: "Rekor loginfo command",
	Long:  `Prints info about the transparency log`,
	Run: func(cmd *cobra.Command, args []string) {
		log := log.Logger
		rekorClient, err := GetRekorClient(viper.GetString("rekor_server"))
		if err != nil {
			log.Fatal(err)
		}

		result, err := rekorClient.Tlog.GetLogInfo(nil)
		if err != nil {
			log.Fatal(err)
		}

		logInfo := result.GetPayload()
		fmt.Printf("Tree Size: %v, Root Hash: %v\n", *logInfo.TreeSize, *logInfo.RootHash)
	},
}

func init() {
	rootCmd.AddCommand(logInfoCmd)
}
