package passwordgenerator

type IPasswordGeneratorService interface {
	Generate(PasswordOptions) string
}

type PasswordOptions struct {
	Lower   int
	Upper   int
	Numbers int
	Special int
}
