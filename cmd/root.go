package cmd

import (
	"fmt"
	"os"

	"github.com/sparsh2/pmgr/cmd/add"
	"github.com/sparsh2/pmgr/cmd/copy"
	"github.com/sparsh2/pmgr/cmd/list"
	"github.com/sparsh2/pmgr/cmd/remove"
	"github.com/sparsh2/pmgr/cmd/update"
	"github.com/sparsh2/pmgr/common"
	passwordmanager "github.com/sparsh2/pmgr/password_manager"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     common.RootCommand,
	Short:   fmt.Sprintf("%v is a password manager", common.RootCommand),
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(copy.CopyCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(update.UpdateCmd)

	passwordmanager.PswManager = passwordmanager.NewPasswordManager(common.SecretKey)
}
