package passwordmanager

import passwordgenerator "github.com/sparsh2/pmgr/password_manager/services/password_generator"

type IPasswordManager interface {
	Add(string, string, passwordgenerator.PasswordOptions) error
	Copy(string) error
	Remove(string) error
	List() ([]string, error)
	Update(string, string) error
}
