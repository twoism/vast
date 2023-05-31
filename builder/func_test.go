package builder

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"go/token"
	"os"
	"testing"
)

func TestNewFunc(t *testing.T) {
	f := NewFunc("hello")
	assert.Equal(t, "hello", f.Name.Name)
}

func TestAddArg(t *testing.T) {
	fn := NewFunc("hello").AddArg(
		NewField("name", "string")).
		AddResults(
			NewField("", "string"))
	assert.Equal(t, "hello", fn.Name.Name)

	fn.AddReturnStatement(&ast.BasicLit{
		Kind:  token.STRING,
		Value: `"Hello, " + name`,
	})

	f := NewFile("test").AddFunc(fn)
	assert.Equal(t, 1, len(f.Funcs()))
	assert.NoError(t, f.Print(os.Stdout))
}
