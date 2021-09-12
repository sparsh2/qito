package copy

import (
	"fmt"

	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	"github.com/spf13/cobra"
)

var CopyCmd = &cobra.Command{
	Use:   common.CopyCommand,
	Short: fmt.Sprintf("%v is used to copy password for a given key", common.CopyCommand),
	RunE:  copyCommand,
}

func copyCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("insufficient arguments")
	}
	err := passwordmanager.PswManager.Copy(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("copied %v password to clipboard!\n", args[0])

	return nil
}
