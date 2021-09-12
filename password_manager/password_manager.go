package passwordmanager

import (
	"github.com/sparsh2/pmgr/common"
	"github.com/sparsh2/pmgr/password_manager/services/clipboard"
	passwordgenerator "github.com/sparsh2/pmgr/password_manager/services/password_generator"
	"github.com/sparsh2/pmgr/password_manager/services/storage"
)

var PswManager IPasswordManager

type PasswordManager struct {
	StorageService    storage.IStorageService
	ClipboardService  clipboard.IClipboardService
	PasswordGenerator passwordgenerator.IPasswordGeneratorService
}

func NewPasswordManager(passpharse string) *PasswordManager {
	return &PasswordManager{
		StorageService:    storage.NewStorageService(passpharse, common.SecretFilepath),
		ClipboardService:  clipboard.NewClipboardService(),
		PasswordGenerator: passwordgenerator.NewPasswordGeneratorService(),
	}
}

func (p *PasswordManager) Add(entity string, password string, options passwordgenerator.PasswordOptions) error {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return err
	}

	if password == "" {
		password = p.PasswordGenerator.Generate(options)
	}

	passwordList[entity] = password

	err = p.StorageService.Save(passwordList)
	if err != nil {
		return err
	}

	return nil
}

func (p *PasswordManager) Copy(entity string) error {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return err
	}

	err = p.ClipboardService.Copy(passwordList[entity].(string))
	if err != nil {
		return err
	}

	return nil
}

func (p *PasswordManager) List() ([]string, error) {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return nil, err
	}

	entityList := []string{}

	for key := range passwordList {
		entityList = append(entityList, key)
	}

	return entityList, err
}

func (p *PasswordManager) Remove(entity string) error {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return err
	}

	delete(passwordList, entity)

	err = p.StorageService.Save(passwordList)
	if err != nil {
		return err
	}

	return nil
}

func (p *PasswordManager) Update(entity string, password string) error {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return err
	}

	passwordList[entity] = password

	err = p.StorageService.Save(passwordList)
	if err != nil {
		return err
	}

	return nil
}
