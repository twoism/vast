package builder

import (
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructFields(t *testing.T) {
	s := NewStruct("S").
		AddFields(
			NewField("A", "string"),
			NewField("B", "string"),
		)
	assert.Equal(t, 2, len(s.Fields.List))
	s.RemoveField("A")
	assert.Equal(t, 1, len(s.Fields.List))
}

func TestToMessageBuilder(t *testing.T) {
	//o := NewStruct("O").
	//	AddFields(
	//		NewField("A", "string"),
	//		NewField("B", "string"),
	//	)
	s := NewStruct("S").
		AddFields(
			NewField("A", "string"),
			NewField("B", "string"),
		)
	m := s.ToProtoBuilder()
	assert.Equal(t, "A", m.GetField("A").GetName())

	d, err := m.Build()
	assert.NoError(t, err)
	//assert.Equal(t, descriptorpb.FieldDescriptorProto_TYPE_MESSAGE,
	//	d.FindFieldByName("S").GetType())

	pp := protoprint.Printer{}
	str, err := pp.PrintProtoToString(d.GetFile())
	assert.NoError(t, err)
	println(str)
}
