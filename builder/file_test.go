package builder

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuilder(t *testing.T) {
	oo := NewStruct("X").AddStringField("Y").
		AddField(NewPointerSelectorField("Z", "time", "Time"))

	f := NewFile("test").AddStructs(
		NewStruct("Person").
			AddStringField("Name").AddStructField("Address", oo),
		NewStruct("Address").
			AddSelectorField("Date", "time", "Time"),
	).AddStruct(oo)

	assert.NoError(t, f.Print(os.Stdout))
}
