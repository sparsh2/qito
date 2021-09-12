package parser

import (
	"encoding/json"
	"fmt"
)

type ParserServiceJSON struct {
}

func GetNewParserServiceJSON() ParserServiceJSON {
	return ParserServiceJSON{}
}

func (p *ParserServiceJSON) Parse(bytes []byte) (map[string]interface{}, error) {
	var mp map[string]interface{}

	if err := json.Unmarshal(bytes, &mp); err != nil {
		return nil, fmt.Errorf("could not parse the file: %v", err.Error())
	}

	return mp, nil
}

func (p *ParserServiceJSON) Unparse(mp map[string]interface{}) ([]byte, error) {
	bytes, err := json.Marshal(mp)

	if err != nil {
		return nil, fmt.Errorf("could not unparse: %v", bytes)
	}

	return bytes, nil
}
