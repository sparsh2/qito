package add

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   common.AddCommand,
	Short: fmt.Sprintf("%v is used to add a key value pair to the current profile", common.AddCommand),
	Run:   addCommand,
}

func addCommand(cmd *cobra.Command, args []string) {
	fmt.Println(args)
}
