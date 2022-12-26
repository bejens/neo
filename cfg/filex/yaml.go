package filex

import (
	"gopkg.in/yaml.v3"
	"io"
)

type YamlParser struct{}

func (yp *YamlParser) Parse(reader io.Reader) (m map[string]any, err error) {

	bs, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bs, m)
	return
}
