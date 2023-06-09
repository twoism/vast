package builder

import (
	"go/ast"
	"go/token"
)

type Const struct {
	*ast.ValueSpec
}

// NewConst creates a new const with the given name and type.
func NewConst(name, fieldType string) *Const {
	return &Const{
		ValueSpec: &ast.ValueSpec{
			Names: []*ast.Ident{ast.NewIdent(name)},
			Type:  ast.NewIdent(fieldType),
		},
	}
}

// ToDecl converts the const to a declaration.
func (c *Const) ToDecl() ast.Decl {
	return &ast.GenDecl{
		Tok:   token.CONST,
		Specs: []ast.Spec{c.ValueSpec},
	}
}
