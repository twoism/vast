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

// NewFromAstField creates a new field from an ast.Field.
func NewFromAstField(field *ast.Field) *Field {
	return &Field{
		Field: field,
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

func NewStructField(name string, s *Struct, pkg string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(pkg),
				Sel: ast.NewIdent(s.Name),
			},
		},
	}
}

// FieldsFromAstFields creates a slice of fields from a slice of ast.Fields.
func FieldsFromAstFields(fields []*ast.Field) []*Field {
	f := make([]*Field, len(fields))

	for i, field := range fields {
		f[i] = NewFromAstField(field)
	}

	return f
}

// FieldName returns the name of the field.
func (f *Field) FieldName() string {
	return f.Names[0].Name
}

// FieldType returns the type of the field.
func (f *Field) FieldType() string {
	return f.Field.Type.(*ast.Ident).Name
}
