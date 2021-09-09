package remove

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   common.RemoveCommand,
	Short: fmt.Sprintf("%v is used to remove password to clipboard", common.RemoveCommand),
	Run:   removeCommand,
}

func removeCommand(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
