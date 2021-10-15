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
	"io/ioutil"
	"os"
	"thanos/app/run"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ArgsApply struct {
	Environment string
	Tag         string
}

var argsApply ArgsApply

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply definitions",
	Long:  `Use to apply definitions kubernetes server`,
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		b, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("Error to read file", file)
			os.Exit(1)
		}
		run.ApplyRun(viper.GetString("THANOS_ADDRESS"), viper.GetString("THANOS_SECRET"), argsApply.Environment, argsApply.Tag, b)
	},
}

func init() {
	applyCmd.Flags().StringVar(&argsApply.Environment, "environment", "prod", "Environment Apply")
	applyCmd.Flags().StringVar(&argsApply.Tag, "tag", "latest", "Image TAG")
	rootCmd.AddCommand(applyCmd)
}
