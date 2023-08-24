package builder

import (
	prb "github.com/jhump/protoreflect/desc/builder"
	"go/ast"
)

type Field struct {
	*ast.Field
}

type FieldOpt func(*Field)

func IsPointer() FieldOpt {
	return func(f *Field) {
		f.Field.Type = &ast.StarExpr{
			X: f.Field.Type,
		}
	}
}

// NewField creates a new field with the given name and type.
func NewField(name string, fieldType string, opts ...FieldOpt) *Field {
	f := &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type:  ast.NewIdent(fieldType),
		},
	}

	for _, opt := range opts {
		opt(f)
	}

	return f
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

func NewStructField(name string, pkg string) *Field {
	return &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type: &ast.SelectorExpr{
				X:   ast.NewIdent(pkg),
				Sel: ast.NewIdent(name),
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
	var name string
	switch t := f.Field.Type.(type) {
	case *ast.Ident:
		name = t.Name
	case *ast.StarExpr:
		name = t.X.(*ast.Ident).Name
	case *ast.SelectorExpr:
		name = t.X.(*ast.Ident).Name + "." + t.Sel.Name
	}

	return name
}

// FieldIsPointer returns true if the field is a pointer.
func (f *Field) FieldIsPointer() bool {
	_, ok := f.Field.Type.(*ast.StarExpr)

	return ok
}

// FieldIsSelector returns true if the field is a selector.
func (f *Field) FieldIsSelector() bool {
	_, ok := f.Field.Type.(*ast.SelectorExpr)

	return ok
}

// DescType returns a prb.FieldType for the field type.
func (f *Field) DescType() *prb.FieldType {
	switch f.FieldType() {
	case "int":
		return prb.FieldTypeInt32()
	case "int32":
		return prb.FieldTypeInt32()
	case "int64":
		return prb.FieldTypeInt64()
	case "uint":
		return prb.FieldTypeUInt32()
	case "uint32":
		return prb.FieldTypeUInt32()
	case "uint64":
		return prb.FieldTypeUInt64()
	case "float32":
		return prb.FieldTypeFloat()
	case "float64":
		return prb.FieldTypeDouble()
	case "bool":
		return prb.FieldTypeBool()
	case "string":
		return prb.FieldTypeString()
	case "[]byte":
		return prb.FieldTypeBytes()
	default:
		msg := prb.NewMessage(f.FieldName())
		return prb.FieldTypeMessage(msg)
	}
}
