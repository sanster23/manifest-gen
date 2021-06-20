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
	cmd := &cobra.Command{
		Use:   "create argo manifest",
		Short: "create a new argo manifest",
		Long:  "",
	}
	return cmd
}
