package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "1.0.0"
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version describes the current version of the pmgr",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("pmgr version %v\n", version)
	},
}
