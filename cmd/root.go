package cmd

import (
	"fmt"
	"os"

	"github.com/sparsh2/qito/cmd/add"
	"github.com/sparsh2/qito/cmd/copy"
	"github.com/sparsh2/qito/cmd/list"
	"github.com/sparsh2/qito/cmd/remove"
	"github.com/sparsh2/qito/cmd/update"
	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     common.RootCommand,
	Short:   fmt.Sprintf("%v is a password manager", common.RootCommand),
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		// fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(add.AddCmd)
	rootCmd.AddCommand(copy.CopyCmd)
	rootCmd.AddCommand(list.ListCmd)
	rootCmd.AddCommand(remove.RemoveCmd)
	rootCmd.AddCommand(update.UpdateCmd)

	passwordmanager.PswManager = passwordmanager.NewPasswordManager(common.SecretKey)
}
