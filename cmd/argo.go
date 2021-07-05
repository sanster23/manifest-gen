package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newArgoCommand())
}

type argoOptions struct {
	appName  string
	gitRepo  string
	autoSync bool
}

func newArgoCommand() *cobra.Command {
	o := argoOptions{}

	var cmd = &cobra.Command{
		Use:   "argo",
		Short: "Genrate argo application manifest",
		Long:  `Generate argo application manifest with a predefined template`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("argo called")
		},
	}

	f := cmd.Flags()
	f.StringVar(&o.appName, "app-name", "hello-world", "app name for argo manifest")
	f.StringVar(&o.gitRepo, "git-repo", "", "git repo link")
	f.BoolVar(&o.autoSync, "auto-sync", false, "enable auto sync on in argo application")

	return cmd

}
