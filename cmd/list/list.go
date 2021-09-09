package list

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   common.ListCommand,
	Short: fmt.Sprintf("%v is used to list different keys", common.ListCommand),
	Run:   listCommand,
}

func listCommand(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
