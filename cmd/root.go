package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	glb "gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg"
	"gitlab.com/heb-engineering/teams/spm-eng/appcloud/tools/gitlab-buddy/pkg/config"
	"go.uber.org/zap"
)

var cfgFilePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitlab-buddy",
	Short: "Management tool for Gitlab content",
	Long: `
	
A management tool for Gitlab intended to reduce toil and automate tasks.
'gitlab-buddy' will serve as a "catch-all" Gitlab automation tool. Although limited now,
is was writting as an extensible framework for future development of complex tasks to integrate
into automations and build flows.

=====================================================================
|| NOTICE: Functionality is currently limited to branch migration. ||
|| Have feature suggestions for this tool? Let the AppCloud team   ||
|| know! Slack us as #appcloud-engcore                             ||
=====================================================================

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFilePath, "config", "", "config file (default is $HOME/.heb-digital/gitlab-buddy.yaml)")
}

// initConfig reads in config file
// This should be called at the sub-command level
// some sub commands may not require a config file
func initConfig() {
	// Check if config flag was defined
	if cfgFilePath != "" {
		cf, err := config.NewConfigFile(cfgFilePath)
		if err != nil {
			return
		}

		err = config.LoadFile(cf)
		if err != nil {
			return
		}
	} else {
		if exists, _ := config.ConfigExists(); !exists {
			err := configInitWizard()
			if err != nil {
				return
			}
		}
		// TODO: Move this into main.go, so that CLI isn't the only easy entry point
		// Attempt to load config file and/or defaults
		// Don't panic if error to allow all defaults to get through
		err := config.Load()
		if err != nil {
			fmt.Errorf("Error loading config: %w", err)
		}

		// Initialize global logger
		glb.InitLogger()
		zap.S().Debug("replaced zap's global loggers")
	}
}
