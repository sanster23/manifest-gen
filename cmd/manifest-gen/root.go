package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var globalUsage = `

manifest-gen is a template generator and template renderer for argo and helm charts
`

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "manifest-gen",
		Short:        "The manifest generator for helm and argo.",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	flags := cmd.PersistentFlags()

	// We can safely ignore any errors that flags.Parse encounters since
	// those errors will be caught later during the call to cmd.Execution.
	// This call is required to gather configuration information prior to
	// execution.
	flags.ParseErrorsWhitelist.UnknownFlags = true
	if err := flags.Parse(args); err != nil {
		fmt.Printf("%+v", err)
	}

	// Add subcommands
	cmd.AddCommand(
		// chart commands
		newCreateCmd(out),

		// Hidden documentation generator command: 'helm docs'
		newDocsCmd(out),
	)

	return cmd, nil
}
