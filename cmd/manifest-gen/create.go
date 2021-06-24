package main

import (
	"io"

	"github.com/spf13/cobra"
)

const createDesc = `
Create provides ability generate manifest for argo and helm charts.
`

func newCreateCmd(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [keyword]",
		Short: "create a new chart with the given name",
		Long:  createDesc,
	}

	cmd.AddCommand(newCreateArgoCmd(out))
	cmd.AddCommand(newCreateHelmCmd(out))
	return cmd
}
