package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate manifest-gen input configurations",
	Long: `Validate command can be used to validate the config file.

	For example:
	
		manifest-gen config validate --config .manifest-gen.json
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("validate called")
	},
}

func init() {
	configCmd.AddCommand(validateCmd)
}
