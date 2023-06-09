package builder

import (
	"go/ast"
	"go/token"
)

type Var struct {
	*ast.ValueSpec
}

// NewVar creates a new var with the given name and type.
func NewVar(name, fieldType string) *Var {
	return &Var{
		ValueSpec: &ast.ValueSpec{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type:  ast.NewIdent(fieldType),
		},
	}
}

// ToDecl converts the var to a declaration.
func (v *Var) ToDecl() ast.Decl {
	return &ast.GenDecl{
		Tok:   token.VAR,
		Specs: []ast.Spec{v.ValueSpec},
	}
}
