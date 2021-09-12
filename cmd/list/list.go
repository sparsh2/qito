package list

import (
	"fmt"

	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   common.ListCommand,
	Short: fmt.Sprintf("%v is used to list different keys", common.ListCommand),
	RunE:  listCommand,
}

func listCommand(cmd *cobra.Command, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("insufficient arguments")
	}
	list, err := passwordmanager.PswManager.List()
	if err != nil {
		return err
	}

	for _, key := range list {
		fmt.Println(key)
	}

	return nil
}
