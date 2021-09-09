package update

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   common.UpdateCommand,
	Short: fmt.Sprintf("%v is used to update a key-password pair", common.UpdateCommand),
	Run:   updateCommand,
}

func updateCommand(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
