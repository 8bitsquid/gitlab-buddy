/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/internal/config"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := configInitWizard()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	configCmd.AddCommand(initCmd)
}

func configInitWizard() error {
	i, _, _ := promptAddWhatClient()
		if i == 0 {
			host, err := promptAddGitlabClient()
			if err != nil {
				return err
			}

			viper.Set("hosts." + host["name"] + ".token", host["token"])
			fileName := strings.Join([]string{config.CONFIG_FILE_NAME, config.CONFIG_FILE_EXT}, ".")
			configPath := filepath.Join(config.ConfigDir, fileName)

			cfg, err := config.NewConfigFile(configPath)
			if err != nil {
				return err
			}
			
			err = config.CreateConfig(cfg)
			if err != nil {
				return err
			}

			fmt.Println(`
###  Gitlab Buddy is configured and ready to go!
###  Type 'glb --help' to see what Gitlab Buddy can do.
		`)
		}

		return nil
}

