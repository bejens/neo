package parser

type Parser interface {
	Parse() (map[string]any, error)
}
