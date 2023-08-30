package proto

import (
	prb "github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"github.com/twoism/vast/builder"
)

func NewBuilderForFile(f *builder.File) *prb.FileBuilder {
	file := prb.NewFile("howdy.proto")
	file.SetProto3(true)

	for _, s := range f.Structs() {
		msg := prb.NewMessage(s.Name)
		for _, sf := range s.StructFields() {
			field := prb.NewField(sf.Name(), FieldTypeForField(sf))
			msg = msg.AddField(field)
		}
		file = file.AddMessage(msg)
	}

	return file
}

// FieldTypeForField returns a prb.FieldType for the field type.
func FieldTypeForField(f *builder.Field) *prb.FieldType {
	switch f.FieldType() {
	case "int":
		return prb.FieldTypeInt32()
	case "int32":
		return prb.FieldTypeInt32()
	case "int64":
		return prb.FieldTypeInt64()
	case "uint":
		return prb.FieldTypeUInt32()
	case "uint32":
		return prb.FieldTypeUInt32()
	case "uint64":
		return prb.FieldTypeUInt64()
	case "float32":
		return prb.FieldTypeFloat()
	case "float64":
		return prb.FieldTypeDouble()
	case "bool":
		return prb.FieldTypeBool()
	case "string":
		return prb.FieldTypeString()
	case "[]byte":
		return prb.FieldTypeBytes()
	default:
		msg := prb.NewMessage(f.Name())
		return prb.FieldTypeMessage(msg)
	}
}

func PrintFile(f *builder.File) (string, error) {
	pb := NewBuilderForFile(f)
	desc, err := pb.BuildDescriptor()

	if err != nil {
		return "", err
	}

	pp := protoprint.Printer{}
	return pp.PrintProtoToString(desc)
}
