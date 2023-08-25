package builder

import (
	prb "github.com/jhump/protoreflect/desc/builder"
	"go/ast"
	"strings"
)

type Field struct {
	*ast.Field

	isPointer   bool
	hasPackage  bool
	packageName string
}

type FieldOpt func(*Field)

func FieldIsPointerOpt() FieldOpt {
	return func(f *Field) {
		f.isPointer = true

		f.Field.Type = &ast.StarExpr{
			X: f.Field.Type,
		}
	}
}

// NewField creates a new field with the given name and type.
func NewField(name string, ftype string, opts ...FieldOpt) *Field {
	pkg, fieldType, hasPkg := SplitPackage(ftype)

	field := &Field{
		Field: &ast.Field{
			Names: []*ast.Ident{ast.NewIdent(name)},
		},
	}

	if hasPkg {
		field.hasPackage = true
		field.packageName = pkg
		field.Field.Type = &ast.SelectorExpr{
			X:   ast.NewIdent(pkg),
			Sel: ast.NewIdent(fieldType),
		}
	} else {
		field.Field.Type = ast.NewIdent(fieldType)
	}

	for _, opt := range opts {
		opt(field)
	}

	return field
}

// NewFromAstField creates a new field from an ast.Field.
func NewFromAstField(field *ast.Field) *Field {
	return &Field{
		Field: field,
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

// Name returns the name of the field.
func (f *Field) Name() string {
	return f.Names[0].Name
}

// FieldType returns the type of the field.
func (f *Field) FieldType() string {
	var name string
	switch t := f.Field.Type.(type) {
	case *ast.Ident:
		name = t.Name
	case *ast.StarExpr:
		switch xt := t.X.(type) {
		case *ast.SelectorExpr:
			name = xt.X.(*ast.Ident).Name + "." + xt.Sel.Name
		case *ast.Ident:
			name = xt.Name
		}
	case *ast.SelectorExpr:
		name = t.X.(*ast.Ident).Name + "." + t.Sel.Name
	}

	return name
}

// IsPointer returns true if the field is a pointer.
func (f *Field) IsPointer() bool {
	return f.isPointer
}

// HasPackage returns true if the field is a selector.
func (f *Field) HasPackage() bool {
	return f.hasPackage
}

// PackageName returns the package name of the field.
func (f *Field) PackageName() string {
	return f.packageName
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
		msg := prb.NewMessage(f.Name())
		return prb.FieldTypeMessage(msg)
	}
}

func SplitPackage(name string) (pkg, fieldType string, hasPkg bool) {
	parts := strings.Split(name, ".")
	if len(parts) == 2 {
		return parts[0], parts[1], true
	}

	return "", name, false
}
