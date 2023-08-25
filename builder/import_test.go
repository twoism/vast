package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewImport(t *testing.T) {
	t.Run("Standard Library", func(t *testing.T) {
		i := NewImport("time")
		assert.Equal(t, "time", i.Path())
		assert.Equal(t, "time", i.BaseName())
		assert.False(t, i.HasAlias())
	})

	t.Run("Standard Library with Alias", func(t *testing.T) {
		i := NewImport("time", ImportAliasOpt("t"))
		assert.Equal(t, "time", i.Path())
		assert.Equal(t, "time", i.BaseName())
		assert.Equal(t, "t", i.Alias())
		assert.True(t, i.HasAlias())
	})

	t.Run("Third Party", func(t *testing.T) {
		i := NewImport("github.com/foo/3p")
		assert.Equal(t, "github.com/foo/3p", i.Path())
		assert.Equal(t, "3p", i.BaseName())
		assert.Equal(t, "", i.Alias())
		assert.False(t, i.HasAlias())
	})

	t.Run("Third Party with Alias", func(t *testing.T) {
		i := NewImport("github.com/foo/3p", ImportAliasOpt("t"))
		assert.Equal(t, "github.com/foo/3p", i.Path())
		assert.Equal(t, "3p", i.BaseName())
		assert.Equal(t, "t", i.Alias())
		assert.True(t, i.HasAlias())
	})
}
