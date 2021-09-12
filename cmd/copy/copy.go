package copy

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	passwordmanager "github.com/sparsh2/pmgr/password_manager"
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

	return nil
}
