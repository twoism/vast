package builder

import (
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
