package storage

type IStorageService interface {
	Read() (map[string]interface{}, error)
	Save(map[string]interface{}) error
	SetFilepath(string) error
}
