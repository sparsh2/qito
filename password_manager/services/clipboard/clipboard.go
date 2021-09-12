package clipboard

import "github.com/atotto/clipboard"

type ClipboardService struct {
}

func NewClipboardService() *ClipboardService {
	return &ClipboardService{}
}

func (c *ClipboardService) Copy(str string) error {
	err := clipboard.WriteAll(str)
	if err != nil {
		return err
	}

	return nil
}
