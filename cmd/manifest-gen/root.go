package main

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
)

var globalUsage = `

manifest-gen is a template generator and template renderer for argo and helm charts
`

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "manifest-gen",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func newRootCmd(out io.Writer, args []string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "manifest-gen",
		Short:        "The manifest generator for helm and argo.",
		Long:         globalUsage,
		SilenceUsage: true,
	}
	flags := cmd.PersistentFlags()

	// addKlogFlags(flags)

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
		// newCreateCmd(out),

		// Hidden documentation generator command: 'helm docs'
		newDocsCmd(out),
	)

	return cmd, nil
}
