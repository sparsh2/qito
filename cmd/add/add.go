package add

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	passwordmanager "github.com/sparsh2/pmgr/password_manager"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   common.AddCommand,
	Short: fmt.Sprintf("%v is used to add a key value pair to the current profile", common.AddCommand),
	RunE:  addCommand,
}

func addCommand(c *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("insufficient arguments")
	}
	err := passwordmanager.PswManager.Add(args[0], args[1])
	if err != nil {
		return err
	}

	return nil
}
