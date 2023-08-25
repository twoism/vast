package builder

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/descriptorpb"
	"os"
	"testing"
)

func printSource(t *testing.T, field *Field) {
	file := NewFile("test").AddStructs(
		NewStruct("Person").AddFields(field),
	)
	assert.NoError(t, file.Print(os.Stdout))
}

func TestNewField(t *testing.T) {

	t.Run("NewField - Value", func(t *testing.T) {
		field := NewField("Age", "int")
		assert.Equal(t, "Age", field.Name())
		assert.Equal(t, "int", field.FieldType())
		assert.False(t, field.IsPointer())
	})

	t.Run("NewField - Pointer", func(t *testing.T) {
		field := NewField("Age", "int", FieldIsPointerOpt())
		assert.Equal(t, "Age", field.Name())
		assert.Equal(t, "int", field.FieldType())
		assert.True(t, field.IsPointer())
	})

	t.Run("NewField - Package Pointer", func(t *testing.T) {
		field := NewField("Date", "time.Time", FieldIsPointerOpt())
		assert.Equal(t, "Date", field.Name())
		assert.Equal(t, "time.Time", field.FieldType())
		assert.Equal(t, "time", field.PackageName())
		assert.True(t, field.IsPointer())
	})

	t.Run("NewField - Package", func(t *testing.T) {
		field := NewField("Date", "time.Time")
		assert.Equal(t, "Date", field.Name())
		assert.Equal(t, "time.Time", field.FieldType())
		assert.False(t, field.IsPointer())
	})

}

func TestDescType(t *testing.T) {
	f := NewField("i", "int")
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_INT32, f.DescType().GetType())

	f = NewStructField("thing", "test")
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE, f.DescType().GetType())
}
