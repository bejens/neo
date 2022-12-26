package parser

import "io"

type Parser interface {
	Parse(reader io.Reader) (map[string]any, error)
}
