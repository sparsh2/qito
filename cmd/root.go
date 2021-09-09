package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "pmgr",
	Short:   "Pmgr is a password manager tool",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Hello world!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
