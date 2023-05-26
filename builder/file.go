package builder

import (
	"go/ast"
	"go/printer"
	"go/token"
	"io"
)

type File struct {
	*ast.File
}

func NewFile(pkg string) *File {
	return &File{File: &ast.File{
		Name: ast.NewIdent(pkg),
	}}
}

// AddStructs adds structs to the file.
func (f *File) AddStructs(structs ...*Struct) *File {
	for _, s := range structs {
		f.AddStruct(s)
	}

	return f
}

// AddStruct adds a struct to the file.
func (f *File) AddStruct(s *Struct) *File {
	f.Decls = append(f.Decls, s.ToDecl())

	return f
}

// Print prints the file to the io.Writer.
func (f *File) Print(w io.Writer) error {
	return printer.Fprint(w, token.NewFileSet(), f.File)
}
