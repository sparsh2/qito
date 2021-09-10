package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
)

// currently, this doesn't work.
// TODO: Migrate to AES encryption

type AESEncryptionService struct {
	aesCipher cipher.Block
}

func (e *AESEncryptionService) setPassphrase(passphrase string) error {
	// get 32 bytes key by hashing passphrase
	key := sha256.Sum256([]byte(passphrase))
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return err
	}

	e.aesCipher = c
	return nil
}

func (e *AESEncryptionService) encrypt(bytes []byte) ([]byte, error) {
	if e.aesCipher == nil {
		return nil, fmt.Errorf("AES cipher not set")
	}

	out := make([]byte, len(bytes))
	e.aesCipher.Encrypt(out, bytes)
	return out, nil
}

func (e *AESEncryptionService) decrypt(bytes []byte) ([]byte, error) {
	if e.aesCipher == nil {
		return nil, fmt.Errorf("AES cipher not set")
	}

	out := make([]byte, len(bytes))
	e.aesCipher.Decrypt(out, bytes)
	return out, nil
}
