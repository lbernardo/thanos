/*
Copyright © 2021 Lucas Bernardo

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
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigClientInput struct {
	Ip     string
	Secret string
}

var configClientInput ConfigClientInput

// config:clientCmd represents the config:client command
var config_clientCmd = &cobra.Command{
	Use:   "config:client",
	Short: "Configure client",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("thanos_secret", configClientInput.Secret)
		viper.Set("thanos_address", configClientInput.Ip)
		viper.WriteConfig()
		fmt.Println("Configured")
	},
}

func init() {
	rootCmd.AddCommand(config_clientCmd)
	config_clientCmd.Flags().StringVar(&configClientInput.Secret, "secret", "", "Secret server")
	config_clientCmd.Flags().StringVar(&configClientInput.Ip, "server", "", "Server ip")
	config_clientCmd.MarkPersistentFlagRequired("secret")
	config_clientCmd.MarkPersistentFlagRequired("server")
}
