package builder

import (
	"go/ast"
	"go/token"
)

type Struct struct {
	*ast.StructType

	Name string
}

// NewStruct creates a new struct with the given name.
func NewStruct(name string) *Struct {
	return &Struct{
		Name: name,
		StructType: &ast.StructType{
			Fields: &ast.FieldList{},
		},
	}
}

// AddSelectorField adds a selector field to the struct.
func (s *Struct) AddSelectorField(name, pkg, fieldType string) *Struct {
	return s.AddField(NewSelectorField(name, pkg, fieldType))
}

// AddStructField adds a struct field to the struct.
func (s *Struct) AddStructField(name string, str *Struct, pkg string) *Struct {
	return s.AddField(NewStructField(name, str, pkg))
}

// AddStringField adds a string field to the struct.
func (s *Struct) AddStringField(name string) *Struct {
	return s.AddField(NewField(name, "string"))
}

// AddField adds a field to the struct.
func (s *Struct) AddField(field *Field) *Struct {
	s.Fields.List = append(s.Fields.List, field.Field)

	return s
}

// AddFields adds fields to the struct.
func (s *Struct) AddFields(fields ...*Field) *Struct {
	for _, field := range fields {
		s.AddField(field)
	}

	return s
}

// RemoveField removes a field from the struct.
func (s *Struct) RemoveField(name string) *Struct {
	for i, field := range s.Fields.List {
		if field.Names[0].Name == name {
			s.Fields.List = append(s.Fields.List[:i], s.Fields.List[i+1:]...)
			break
		}
	}

	return s
}

// ToSpec returns a *ast.TypeSpec for the struct.
func (s *Struct) ToSpec() *ast.TypeSpec {
	return &ast.TypeSpec{
		Name: ast.NewIdent(s.Name),
		Type: s.StructType,
	}
}

// ToDecl returns a *ast.GenDecl for the struct.
func (s *Struct) ToDecl() *ast.GenDecl {
	return &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{s.ToSpec()},
	}
}

// String returns the string representation of the struct.
func (s *Struct) String() string {
	return s.Name
}
