package builder

import "go/ast"

type Statement struct {
	*ast.BlockStmt
}

// NewBlockStatement creates a new statement.
func NewBlockStatement(stmts ...ast.Stmt) *Statement {
	return &Statement{
		BlockStmt: &ast.BlockStmt{
			List: stmts,
		},
	}
}
