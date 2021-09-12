package parser

type IParserService interface {
	Parse([]byte) (map[string]interface{}, error)
	Unparse(map[string]interface{}) ([]byte, error)
}
