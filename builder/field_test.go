package builder

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"testing"
)

func TestNewSelectorField(t *testing.T) {
	f := NewSelectorField("Date", "time", "Time")
	assert.Equal(t, "Date", f.Names[0].Name)
	assert.Equal(t, "time.Time",
		f.Type.(*ast.SelectorExpr).X.(*ast.Ident).
			Name+"."+f.Type.(*ast.SelectorExpr).Sel.Name)
}
