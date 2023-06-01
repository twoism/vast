package builder

import (
	"github.com/stretchr/testify/assert"
	"go/ast"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

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

	s := NewStruct("thing").AddFields(f)
	f = NewStructField("thing", s, "test")
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, f.DescType().GetType())
}
