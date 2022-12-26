package flagx

import (
	"io"
	"os"
)

type FlagParser struct {
	prefix string
	seg    string
	sep    string
}

func (fp *FlagParser) Parse(reader io.Reader) (m map[string]any, err error) {

	flags := os.Args[1:]

}
