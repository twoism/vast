package builder

import (
	"go/ast"
	"go/parser"
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

// NewFromSource creates a new file from the given source.
func NewFromSource(src string) (*File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return &File{File: f}, nil
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

// Structs returns the structs in the file.
func (f *File) Structs() []*Struct {
	var structs []*Struct

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {
						structs = append(structs, &Struct{
							StructType: structType,
							Name:       typeSpec.Name.Name,
						})
					}
				}
			}
		}
	}

	return structs
}

// Struct returns the struct with the given name.
func (f *File) Struct(name string) *Struct {
	for _, s := range f.Structs() {
		if s.Name == name {
			return s
		}
	}

	return nil
}

// Funcs returns the functions in the file.
func (f *File) Funcs() []*Func {
	var funcs []*Func

	for _, decl := range f.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			funcs = append(funcs, &Func{FuncDecl: funcDecl})
		}
	}

	return funcs
}

// Func returns the function with the given name.
func (f *File) Func(name string) *Func {
	for _, fn := range f.Funcs() {
		if fn.Name.Name == name {
			return fn
		}
	}

	return nil
}

// AddFunc adds a function to the file.
func (f *File) AddFunc(fn *Func) *File {
	f.Decls = append(f.Decls, fn.FuncDecl)

	return f
}

// Enums returns the enums in the file.
func (f *File) Enums() []*Enum {
	var enums []*Enum

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					if ident, ok := typeSpec.Type.(*ast.Ident); ok {
						if ident.Obj != nil {
							if _, ok := ident.Obj.Decl.(*ast.ValueSpec); ok {
								enums = append(enums, &Enum{
									Name:      typeSpec.Name.Name,
									ValueSpec: ident.Obj.Decl.(*ast.ValueSpec),
								})
							}
						}
					}
				}
			}
		}
	}

	return enums
}

// Print prints the file to the io.Writer.
func (f *File) Print(w io.Writer) error {
	return printer.Fprint(w, token.NewFileSet(), f.File)
}

func (f *File) PackageName() string {
	return f.Name.Name
}
