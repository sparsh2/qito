package storage

import (
	"fmt"
	"os"

	"github.com/sparsh2/pmgr/password_manager/services/storage/encryption"
	"github.com/sparsh2/pmgr/password_manager/services/storage/parser"
)

type StorageService struct {
	EncryptionService encryption.IEncryptionService
	ParserService     parser.IParserService
	Filepath          string
}

func GetNewStorageService(passphrase string) *StorageService {
	return &StorageService{
		EncryptionService: encryption.GetNewEncryptionServiceXOR(passphrase),
	}
}

func (s *StorageService) Read() (map[string]interface{}, error) {
	rawData, err := os.ReadFile(s.Filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err.Error())
	}

	decryptedBytes, err := s.EncryptionService.Decrypt(rawData)
	if err != nil {
		return nil, fmt.Errorf("could not decrypt file: %v", err.Error())
	}

	parsedPasswords, err := s.ParserService.Parse(decryptedBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse decrypted file: %v", err.Error())
	}

	return parsedPasswords, nil
}
