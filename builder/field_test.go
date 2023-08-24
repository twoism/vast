package builder

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

func TestNewField(t *testing.T) {
	f := NewField("i", "int")
	assert.Equal(t, "i", f.Names[0].Name)
	assert.Equal(t, "int", f.Type.(*ast.Ident).Name)

	f2 := NewField("i", "int", IsPointer())
	assert.Equal(t, "i", f2.Names[0].Name)
	assert.Equal(t, "int", f2.Type.(*ast.StarExpr).X.(*ast.Ident).Name)
}

func TestNewSelectorField(t *testing.T) {
	f := NewSelectorField("Date", "time", "Time")
	assert.Equal(t, "Date", f.Names[0].Name)
	assert.Equal(t, "time.Time",
		f.Type.(*ast.SelectorExpr).X.(*ast.Ident).
			Name+"."+f.Type.(*ast.SelectorExpr).Sel.Name)
}

func TestDescType(t *testing.T) {
	f := NewField("i", "int")
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_INT32, f.DescType().GetType())

	f = NewStructField("thing", "test")
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, f.DescType().GetType())
}
