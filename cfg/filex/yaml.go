package filex

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type YamlParser struct {
	path string
}

func (yp *YamlParser) Parse() (m map[string]any, err error) {

	f, err := os.OpenFile(yp.path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return m, err
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bs, m)
	return
}
