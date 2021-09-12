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

func GetNewStorageService(passphrase string, filepath string) *StorageService {
	return &StorageService{
		EncryptionService: encryption.GetNewEncryptionServiceXOR(passphrase),
		ParserService:     parser.GetNewParserServiceJSON(),
		Filepath:          filepath,
	}
}

func (s *StorageService) SetFilepath(filepath string) error {
	s.Filepath = filepath

	return nil
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

func (s *StorageService) Save(data map[string]interface{}) error {
	bytes, err := s.ParserService.Unparse(data)
	if err != nil {
		return fmt.Errorf("could not unparse: %v", err.Error())
	}

	encryptedBytes, err := s.EncryptionService.Encrypt(bytes)
	if err != nil {
		return fmt.Errorf("could not encrypt file contents: %v", err.Error())
	}

	err = os.WriteFile(s.Filepath, encryptedBytes, 0666)
	if err != nil {
		return fmt.Errorf("could not save passwords to file: %v", err.Error())
	}

	return nil
}
