package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFunc(t *testing.T) {
	f := NewFunc("hello")
	assert.Equal(t, "hello", f.Name.Name)
}

func TestAddArg(t *testing.T) {
	f := NewFunc("hello").AddArg(NewField("name", "string")).AddResults(NewField("", "string"))
	assert.Equal(t, "hello", f.Name.Name)
}
