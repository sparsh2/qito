package encryption

type IEncryptionService interface {
	encrypt([]byte) ([]byte, error)
	decrypt([]byte) ([]byte, error)
	setPassphrase(string) error
}
