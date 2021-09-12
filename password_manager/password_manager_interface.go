package passwordmanager

type IPasswordManager interface {
	Add(string, string) error
	Copy(string) error
	Remove(string) error
	List() error
	Update(string, string) error
}
