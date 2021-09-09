package copy

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/spf13/cobra"
)

var CopyCmd = &cobra.Command{
	Use:   common.CopyCommand,
	Short: fmt.Sprintf("%v is used to copy password for a given key", common.CopyCommand),
	Run:   copyCommand,
}

func copyCommand(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
