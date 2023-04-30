package envx

import (
	"os"
	"testing"
)

func TestEnvParser_Parse(t *testing.T) {

	parser := EnvParser{
		Prefix: "neo",
		Seg:    ".",
		Sep:    "=",
	}

	_ = os.Setenv("neo.schema.key1", "1")
	_ = os.Setenv("neo.schema.key2", "2")
	_ = os.Setenv("neo.schema.a.b", "1")
	_ = os.Setenv("neo.schema.a.c", "2")
	_ = os.Setenv("neo.schema.a.e.d[1]", "2")
	_ = os.Setenv("neo.schema.a.e.d[0]", "1")
	_ = os.Setenv("neo.schema.a.e.f[0].a", "a")
	_ = os.Setenv("neo.schema.a.e.f[1].b", "b")

	m, err := parser.Parse()
	if err != nil {
		t.Error(err)
	}

	t.Log(m)
}
