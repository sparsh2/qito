package add

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	passwordmanager "github.com/sparsh2/pmgr/password_manager"
	passwordgenerator "github.com/sparsh2/pmgr/password_manager/services/password_generator"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   common.AddCommand,
	Short: fmt.Sprintf("%v is used to add a key value pair to the current profile", common.AddCommand),
	RunE:  addCommand,
}

func addCommand(c *cobra.Command, args []string) error {
	if len(args) != 2 && len(args) != 1 {
		return fmt.Errorf("insufficient arguments")
	}

	options := passwordgenerator.PasswordOptions{}
	password := ""

	if len(args) == 1 {
		fmt.Println("Enter # of characters of type")
		fmt.Printf("Lowercase:")
		fmt.Scanf("%d", &options.Lower)
		fmt.Printf("Uppercase:")
		fmt.Scanf("%d", &options.Upper)
		fmt.Printf("Numbers:")
		fmt.Scanf("%d", &options.Numbers)
		fmt.Printf("Special:")
		fmt.Scanf("%d", &options.Special)
	} else {
		password = args[1]
	}

	err := passwordmanager.PswManager.Add(args[0], password, options)
	if err != nil {
		return err
	}

	return nil
}
