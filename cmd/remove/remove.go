package remove

import (
	"fmt"

	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   common.RemoveCommand,
	Short: fmt.Sprintf("%v is used to remove password to clipboard", common.RemoveCommand),
	RunE:  removeCommand,
}

func removeCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("insufficient arguments")
	}
	err := passwordmanager.PswManager.Remove(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("deleted %v entry from database\n", args[0])

	return nil
}
