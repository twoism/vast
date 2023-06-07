# vast

Vast is a go/ast manipulation library. It provides a set of builders and helpers to create and modify go/ast nodes along with the ability to render them back to valid go source code. It can also read and write Protobuf descriptors using the [protoreflect](github.com/jhump/protoreflect/desc/builder) library.

## Installation

```bash
go get -u github.com/twoism/vast
```

## Usage

```go
func TestBuilderExample(t *testing.T) {
	fn := NewFunc("hello").
		AddArg(NewField("name", "string")).
		AddResults(NewField("", "string"))

	otherStruct := NewStruct("X").
		AddStringField("Y").
		AddField(NewPointerSelectorField("Z", "time", "Time"))

	file := NewFile("test").AddStructs(
		NewStruct("Person").
			AddStringField("Name").
			AddStructField("Address", otherStruct, "test"),
		NewStruct("Address").
			AddSelectorField("Date", "time", "Time"),
		otherStruct,
	).AddFunc(fn)

	assert.NoError(t, file.Print(os.Stdout))
}
```

Output:

```go
package test

type Person struct {
	Name	string
	Address	test.X
}
type Address struct {
	Date time.Time
}
type X struct {
	Y	string
	Z	*time.Time
}

func hello(name string) string {}
```
