package builder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuilder(t *testing.T) {
	fn := NewFunc("hello").
		AddArg(NewField("name", "string")).
		AddResults(NewField("", "string"))

	other := NewStruct("X").AddStringField("Y").
		AddField(NewPointerSelectorField("Z", "time", "Time"))

	f := NewFile("test").AddStructs(
		NewStruct("Person").
			AddStringField("Name").AddStructField("Address", other, "test"),
		NewStruct("Address").
			AddSelectorField("Date", "time", "Time"),
		other,
	).AddFunc(fn)

	fmt.Printf("%+v\n", f.Structs())
	assert.NoError(t, f.Print(os.Stdout))
}
