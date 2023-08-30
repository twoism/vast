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
		AddField(NewField("Z", "time.Time", FieldIsPointerOpt()))

	f := NewFile("test").AddStructs(
		NewStruct("Person").
			AddStringField("Name").
			AddStructField("Address", other),
		//NewStruct("Address").AddField("Date", "time.Time"),
		other,
	).AddFunc(fn)

	fmt.Printf("%+v\n", f.Structs())
	assert.NoError(t, f.Print(os.Stdout))
}

func TestBuilderExample2(t *testing.T) {
	file := NewFile("test").AddStructs(
		NewStruct("Person").
			AddField(
				NewField("Date", "time.Time",
					FieldIsPointerOpt()),
			),
	)

	assert.NoError(t, file.PrintFormatted(os.Stdout))
}

func TestBuilderExample(t *testing.T) {
	fn := NewFunc("hello").
		AddArg(NewField("name", "string")).
		AddResults(NewField("", "string"))

	otherStruct := NewStruct("X").
		AddStringField("Y").
		AddField(NewField("T", "time.Time", FieldIsPointerOpt()))

	file := NewFile("test").AddStructs(
		NewStruct("Person").
			AddField(NewField("Name", "string", FieldIsPointerOpt())).
			AddStructField("Address", otherStruct),
		NewStruct("Address").
			AddField(NewField("Date", "time.Time")),
		otherStruct,
	).AddFunc(fn)

	assert.NoError(t, file.PrintFormatted(os.Stdout))
	t.TempDir()
}

var src = `
package test

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

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
				NewStructField("Person", f.PackageName()),
			),
	)
	assert.NoError(t, f.Print(os.Stdout))
}

func TestStructFuncs(t *testing.T) {
	f, err := NewFromSource(src)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(f.Structs()))
	fns := f.StructFuncs("Person")
	fmt.Printf("%+v\n", fns[0].Name)
}
