package encryption

import "fmt"

type EncryptionServiceXOR struct {
	passphrase string
}

func (e *EncryptionServiceXOR) setPassphrase(passphrase string) error {
	e.passphrase = passphrase

	return nil
}

func (e *EncryptionServiceXOR) encrypt(bytes []byte) ([]byte, error) {
	if e.passphrase == "" {
		return nil, fmt.Errorf("passphrase not initialized")
	}

	out := make([]byte, len(bytes))
	for i := 0; i < len(out); i++ {
		out[i] = bytes[i] ^ byte(e.passphrase[i%len(e.passphrase)])
	}

	return out, nil
}

func (e *EncryptionServiceXOR) decrypt(bytes []byte) ([]byte, error) {
	if e.passphrase == "" {
		return nil, fmt.Errorf("passphrase not initialized")
	}

	out := make([]byte, len(bytes))
	for i := 0; i < len(out); i++ {
		out[i] = bytes[i] ^ byte(e.passphrase[i%len(e.passphrase)])
	}

	return out, nil
}
