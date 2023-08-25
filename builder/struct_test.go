package builder

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/descriptorpb"
	"testing"
)

func TestStructFields(t *testing.T) {
	t.Run("AddField-Primative", func(t *testing.T) {
		s := NewStruct("S").
			AddField(NewField("A", "string"))
		assert.Equal(t, 1, len(s.Fields.List))
		assert.Equal(t, 1, len(s.Imports))
	})

	//s := NewStruct("S").
	//	AddFields(
	//		NewField("A", "string"),
	//		NewField("B", "time.Time"),
	//		NewField("C", "github.com/foo/bar.Time"),
	//	)
}

func TestToMessageBuilder(t *testing.T) {
	o := NewStruct("O").
		AddFields(
			NewField("A", "string"),
			NewField("B", "string"),
		)
	s := NewStruct("S").
		AddFields(
			NewField("A", "string"),
			NewField("B", "string"),
		).AddStructField("O", o)
	m := s.ToProtoBuilder()
	assert.Equal(t, "A", m.GetField("A").GetName())

	d, err := m.Build()
	assert.NoError(t, err)
	assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
		d.FindFieldByName("O").GetType())
}
