package main

import (
	"io"

	"github.com/spf13/cobra"
)

var createArgoDesc = `
This command creates an argo manifest with a predefined schema and given input
`

type argoOptions struct {
	appName  string
	gitRepo  string
	autoSync bool
}

func newCreateArgoCmd(out io.Writer) *cobra.Command {
	o := &argoOptions{}

	cmd := &cobra.Command{
		Use:   "create argo manifest",
		Short: "create a new argo manifest",
		Long:  createArgoDesc,
	}
	f := cmd.Flags()
	f.StringVar(&o.appName, "app-name", "hello-world", "app name for argo manifest")
	f.StringVar(&o.gitRepo, "git-repo", "", "git repo link")
	f.BoolVar(&o.autoSync, "auto-sync", false, "enable auto sync on in argo application")

	return cmd
}
