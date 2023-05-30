package builder

import "go/ast"

type Func struct {
	*ast.FuncDecl
}

// NewFunc creates a new function with the given name.
func NewFunc(name string) *Func {
	return &Func{
		FuncDecl: &ast.FuncDecl{
			Name: ast.NewIdent(name),
			Type: &ast.FuncType{
				Params:  &ast.FieldList{},
				Results: &ast.FieldList{},
			},
		},
	}
}

// AddArgs adds arguments to the function.
func (f *Func) AddArgs(args ...*Field) *Func {
	for _, arg := range args {
		f.AddArg(arg)
	}

	return f
}

// AddArg adds an argument to the function.
func (f *Func) AddArg(arg *Field) *Func {
	f.Type.Params.List = append(f.Type.Params.List, arg.Field)

	return f
}

// AddResults adds results to the function.
func (f *Func) AddResults(results ...*Field) *Func {
	for _, result := range results {
		f.AddResult(result)
	}

	return f
}

// AddResult adds a result to the function.
func (f *Func) AddResult(result *Field) *Func {
	f.Type.Results.List = append(f.Type.Results.List, result.Field)

	return f
}

// AddBody adds a body to the function.
func (f *Func) AddBody(body ...ast.Stmt) *Func {
	f.Body.List = append(f.Body.List, body...)

	return f
}

// AddReturn adds a return statement to the function.
func (f *Func) AddReturn(exprs ...ast.Expr) *Func {
	f.AddBody(&ast.ReturnStmt{
		Results: exprs,
	})

	return f
}

// AddComment adds a comment to the function.
func (f *Func) AddComment(comment string) *Func {
	f.Doc = &ast.CommentGroup{
		List: []*ast.Comment{
			&ast.Comment{
				Text: "// " + comment,
			},
		},
	}

	return f
}

// AddReceiver adds a receiver to the function.
func (f *Func) AddReceiver(receiver *Field) *Func {
	f.Recv = &ast.FieldList{
		List: []*ast.Field{
			receiver.Field,
		},
	}

	return f
}
