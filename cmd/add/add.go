package add

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/sparsh2/qito/common"
	passwordmanager "github.com/sparsh2/qito/password_manager"
	passwordgenerator "github.com/sparsh2/qito/password_manager/services/password_generator"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   common.AddCommand,
	Short: fmt.Sprintf("%v is used to add a key value pair to the current profile", common.AddCommand),
	RunE:  addCommand,
}

func addCommand(c *cobra.Command, args []string) error {
	if len(args) != 1 && len(args) != 2 {
		return fmt.Errorf("incorrect number of arguments arguments, expected 1 or 2, got %d", len(args))
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
		validate := func(str string) error {
			if str == "" {
				return fmt.Errorf("password cannot be empty")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Password",
			Validate: validate,
			Mask:     '*',
		}

		result, err := prompt.Run()

		if err != nil {
			return fmt.Errorf("prompt failed %v", err)
		}
		password = result
	}

	err := passwordmanager.PswManager.Add(args[0], password, options)
	if err != nil {
		return err
	}

	fmt.Printf("added %v password to database\n", args[0])

	return nil
}
