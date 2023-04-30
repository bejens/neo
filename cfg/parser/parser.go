package parser

type Parser interface {
	Ext() string
	Parse() (map[string]any, error)
}
