package update

import (
	"fmt"

	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   common.UpdateCommand,
	Short: fmt.Sprintf("%v is used to update a key-password pair", common.UpdateCommand),
	RunE:  updateCommand,
}

func updateCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("insufficient arguments")
	}
	err := passwordmanager.PswManager.Update(args[0], args[1])
	if err != nil {
		return err
	}

	return nil
}
