package passwordmanager

import (
	"fmt"

	"github.com/sparsh2/pmgr/common"
	"github.com/sparsh2/pmgr/password_manager/services/clipboard"
	"github.com/sparsh2/pmgr/password_manager/services/storage"
)

var PswManager IPasswordManager

type PasswordManager struct {
	StorageService   storage.IStorageService
	ClipboardService clipboard.IClipboardService
}

func NewPasswordManager(passpharse string) *PasswordManager {
	return &PasswordManager{
		StorageService:   storage.GetNewStorageService(passpharse, common.SecretFilepath),
		ClipboardService: clipboard.NewClipboardService(),
	}
}

func (p *PasswordManager) Add(entity string, password string) error {
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

func (p *PasswordManager) List() error {
	passwordList, err := p.StorageService.Read()
	if err != nil {
		return err
	}

	for key := range passwordList {
		fmt.Println(key)
	}

	return nil
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
