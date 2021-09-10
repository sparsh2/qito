package services

type IEncryptionService interface {
	encrypt([]byte) ([]byte, error)
	decrypt([]byte) ([]byte, error)
}
