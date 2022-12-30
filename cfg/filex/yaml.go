package filex

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/bejens/neo/logx"

	"gopkg.in/yaml.v3"
)

type YamlParser struct {
	Path string
}

func (yp *YamlParser) Parse() (m map[string]any, err error) {

	f, err := os.Open(yp.Path)
	if err != nil {
		if errors.Is(err, fs.ErrExist) {
			logx.Warn("config file is not exist")
			return m, nil
		}
		return m, err
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bs, m)
	return
}
