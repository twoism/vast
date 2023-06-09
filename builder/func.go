package builder

import (
	"go/ast"
)

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
			Body: &ast.BlockStmt{},
		},
	}
}

// Args returns the arguments of the function as a slice of Field.
func (f *Func) Args() []*Field {
	return FieldsFromAstFields(f.Type.Params.List)
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

// AddReturnStatement adds a return statement to the function.
func (f *Func) AddReturnStatement(exprs ...ast.Expr) *Func {
	f.AddBody(&ast.ReturnStmt{
		Results: exprs,
	})

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

// Receiver returns the receiver of the function.
func (f *Func) Receiver() *Field {
	return NewFromAstField(f.Recv.List[0])
}

// Results returns the results of the function as a slice of Field.
func (f *Func) Results() []*Field {
	return FieldsFromAstFields(f.Type.Results.List)
}

// String returns the string representation of the function.
func (f *Func) String() string {
	return f.Name.Name
}
