package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config generate and validate config file",
	Long: `Config command can be used to perform actions over config file for manifest-gen
	
	User can generate a sample config for manifest-gen with a predefined structure and default values.
	User can also validate their configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
