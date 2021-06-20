package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// var settings = cli.New()

func init() {
	log.SetFlags(log.Lshortfile)
}

func warning(format string, v ...interface{}) {
	format = fmt.Sprintf("WARNING: %s\n", format)
	fmt.Fprintf(os.Stderr, format, v...)
}

func main() {

	cmd, err := newRootCmd(os.Stdout, os.Args[1:])
	if err != nil {
		warning("%+v", err)
		os.Exit(1)
	}

	// run when each command's execute method is called
	cobra.OnInitialize(func() {
		fmt.Print("cobra onInit")
	})

	if err := cmd.Execute(); err != nil {
		fmt.Printf("%+v", err)
	}
}
