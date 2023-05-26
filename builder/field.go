package builder

import "go/ast"

type Field struct {
	*ast.Field
}

// NewField creates a new field with the given name and type.
func NewField(name, fieldType string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type:  ast.NewIdent(fieldType),
		},
	}
}

// NewPointerField creates a new field with the given name and type.
func NewPointerField(name, fieldType string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.StarExpr{
				X: ast.NewIdent(fieldType),
			},
		},
	}
}

// NewPointerSelectorField creates a new field with the given name and type.
func NewPointerSelectorField(name, pkg, fieldType string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.StarExpr{
				X: &ast.SelectorExpr{
					X:   ast.NewIdent(pkg),
					Sel: ast.NewIdent(fieldType),
				},
			},
		},
	}
}

func NewSelectorField(name, pkg, fieldType string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(pkg),
				Sel: ast.NewIdent(fieldType),
			},
		},
	}
}

func NewStructField(name string, s *Struct) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent("tttt"),
				Sel: ast.NewIdent(s.Name),
			},
		},
	}
}
