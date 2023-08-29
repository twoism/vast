package proto

import (
	"github.com/stretchr/testify/assert"
	vast "github.com/twoism/vast/builder"
	"testing"
)

func TestPrintToProtoFile(t *testing.T) {
	f := vast.NewFile("test").AddStructs(
		vast.NewStruct("Person").
			AddStringField("Name"),
		vast.NewStruct("Address").
			AddStringField("Street"))

	str, err := PrintFile(f)
	assert.NoError(t, err)
	println(str)
}
