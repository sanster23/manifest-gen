package main

import (
	"fmt"
	"io"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const docsDesc = `
Generate documentation files for Helm.

This command can generate documentation for Helm in the following formats:

- Markdown
- Man pages

It can also generate bash autocompletions.
`

type docsOptions struct {
	dest            string
	docTypeString   string
	topCmd          *cobra.Command
	generateHeaders bool
}

func newDocsCmd(out io.Writer) *cobra.Command {
	o := &docsOptions{}

	cmd := &cobra.Command{
		Use:    "docs",
		Short:  "generate documentation as markdown or man pages",
		Long:   docsDesc,
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			o.topCmd = cmd.Root()
			return o.run(out)
		},
	}

	f := cmd.Flags()
	f.StringVar(&o.dest, "dir", "./", "directory to which documentation is written")
	f.StringVar(&o.docTypeString, "type", "markdown", "the type of documentation to generate (markdown, man, bash)")
	f.BoolVar(&o.generateHeaders, "generate-headers", false, "generate standard headers for markdown files")

	return cmd
}

func (o *docsOptions) run(out io.Writer) error {
	switch o.docTypeString {
	case "markdown", "mdown", "md":
		if o.generateHeaders {
			standardLinks := func(s string) string { return s }

			hdrFunc := func(filename string) string {
				base := filepath.Base(filename)
				name := strings.TrimSuffix(base, path.Ext(base))
				title := strings.Title(strings.Replace(name, "_", " ", -1))
				return fmt.Sprintf("---\ntitle: \"%s\"\n---\n\n", title)
			}

			return doc.GenMarkdownTreeCustom(o.topCmd, o.dest, hdrFunc, standardLinks)
		}
		return doc.GenMarkdownTree(o.topCmd, o.dest)
	case "man":
		manHdr := &doc.GenManHeader{Title: "HELM", Section: "1"}
		return doc.GenManTree(o.topCmd, manHdr, o.dest)
	case "bash":
		return o.topCmd.GenBashCompletionFile(filepath.Join(o.dest, "completions.bash"))
	default:
		return errors.Errorf("unknown doc type %q. Try 'markdown' or 'man'", o.docTypeString)
	}
}
