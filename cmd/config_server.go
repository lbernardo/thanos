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
	"os"
	"os/exec"
	"thanos/app/models"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var portInput string
var ipInput string

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config:server",
	Short: "Configure server",
	Run: func(cmd *cobra.Command, args []string) {
		secret := uuid.NewString()
		viper.Set("thanos_secret", secret)
		viper.Set("thanos_address", fmt.Sprintf("0.0.0.0:%v", portInput))
		viper.WriteConfig()

		traefik := models.NewTraefikService(ipInput)
		traefik.ToFile(".traefik-service.json")

		exec.Command("kubectl", "apply", "-f", ".traefik-service.json").Run()
		os.Remove(".traefik-service.json")
		fmt.Println("> Update service traefik")
		fmt.Println("> Configured in", viper.ConfigFileUsed())
		fmt.Printf("\n\nIn client use:\n\033[1;32mthanos config:client --secret %v --server %v:%v\033[0m\n", secret, ipInput, portInput)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVar(&portInput, "port", "3434", "Port to server use thanos")
	configCmd.Flags().StringVar(&ipInput, "ip", "", "Public ip")
	configCmd.MarkFlagRequired("ip")
}
