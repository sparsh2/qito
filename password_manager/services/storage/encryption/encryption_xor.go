package encryption

import "fmt"

type EncryptionServiceXOR struct {
	Passphrase string
}

func NewEncryptionServiceXOR(passphrase string) *EncryptionServiceXOR {
	return &EncryptionServiceXOR{
		Passphrase: passphrase,
	}
}

func (e *EncryptionServiceXOR) SetPassphrase(passphrase string) error {
	e.Passphrase = passphrase

	return nil
}

func (e *EncryptionServiceXOR) Encrypt(bytes []byte) ([]byte, error) {
	if e.Passphrase == "" {
		return nil, fmt.Errorf("passphrase not initialized")
	}

	out := make([]byte, len(bytes))
	for i := 0; i < len(out); i++ {
		out[i] = bytes[i] ^ byte(e.Passphrase[i%len(e.Passphrase)])
	}

	return out, nil
}

func (e *EncryptionServiceXOR) Decrypt(bytes []byte) ([]byte, error) {
	if e.Passphrase == "" {
		return nil, fmt.Errorf("passphrase not initialized")
	}

	out := make([]byte, len(bytes))
	for i := 0; i < len(out); i++ {
		out[i] = bytes[i] ^ byte(e.Passphrase[i%len(e.Passphrase)])
	}

	return out, nil
}
