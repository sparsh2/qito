package encryption

type IEncryptionService interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	SetPassphrase(string) error
}
