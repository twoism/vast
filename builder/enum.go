package builder

import "go/ast"

type Enum struct {
	*ast.ValueSpec

	Name string
}
