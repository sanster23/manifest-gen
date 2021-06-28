package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const createDesc = `Create can be used to generate a single argo application manifest or a whole helm chart 
with a predefined directory structure and some predefined template`

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create argo or helm manifests",
	Long:  createDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
