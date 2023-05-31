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

	other := NewStruct("X").
		AddStringField("Y").
		AddField(NewPointerSelectorField("Z", "time", "Time"))

	f := NewFile("test").AddStructs(
		NewStruct("Person").
			AddStringField("Name").
			AddStructField("Address", other, "test"),
		NewStruct("Address").
			AddSelectorField("Date", "time", "Time"),
		other,
	).AddFunc(fn)

	fmt.Printf("%+v\n", f.Structs())
	assert.NoError(t, f.Print(os.Stdout))
}

var src = `
package test

type Person struct {}

func main() {}
`

func TestNewFromSource(t *testing.T) {
	f, err := NewFromSource(src)
	assert.NoError(t, err)
	assert.Equal(t, "test", f.Name.Name)
	assert.Equal(t, 1, len(f.Structs()))

	f.AddStruct(
		NewStruct("Address").
			AddField(
				NewStructField("Person", f.Struct("Person"), f.PackageName()),
			),
	)
	assert.NoError(t, f.Print(os.Stdout))
}
