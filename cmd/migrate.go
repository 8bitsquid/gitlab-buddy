package cmd

import (
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate content of Gitlab Projects or Groups.",
	Long: `
Migrate content of Gitlab Projects or Groups 
(NOTE: Currently only branch migration is supported).
	
	`,
	Args: cobra.NoArgs,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
