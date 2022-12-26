package envx

import (
	"bytes"
	"os"
	"testing"
)

func TestEnvParser_Parse(t *testing.T) {

	parser := EnvParser{
		prefix: "neo",
		seg:    ".",
		sep:    "=",
	}

	_ = os.Setenv("neo.schema.key1", "1")
	_ = os.Setenv("neo.schema.key2", "2")
	_ = os.Setenv("neo.schema.a.b", "1")
	_ = os.Setenv("neo.schema.a.c", "2")

	m, err := parser.Parse(bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	t.Log(m)
}