package clipboard

type IClipboardService interface {
	Copy(string) error
}
