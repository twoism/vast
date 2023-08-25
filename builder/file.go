package builder

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/imports"
	"io"
)

type File struct {
	*ast.File

	toImport map[string]struct{}
}

func NewFile(pkg string) *File {
	return &File{
		toImport: make(map[string]struct{}),
		File: &ast.File{
			Name:    ast.NewIdent(pkg),
			Imports: make([]*ast.ImportSpec, 0),
		}}
}

// NewFromFile creates a new file from the given file.
func NewFromFile(filename string) (*File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return &File{File: f}, nil
}

// NewFromSource creates a new file from the given source.
func NewFromSource(src string) (*File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), "", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return &File{File: f}, nil
}

func (f *File) PackageName() string {
	return f.Name.Name
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
	for imp, _ := range s.Imports {
		f.AddImport(imp)
	}

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

// StructFuncs Funcs for the given struct.
func (f *File) StructFuncs(name string) []*Func {
	var funcs []*Func

	for _, fn := range f.Funcs() {
		if fn.Recv != nil {
			if fn.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name == name {
				funcs = append(funcs, fn)
			}
		}
	}

	return funcs
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

// Constants returns the constants in the file.
func (f *File) Constants() []*Const {
	var consts []*Const

	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					consts = append(consts, NewConst(valueSpec.Names[0].Name, valueSpec.Type.(*ast.Ident).Name))
				}
			}
		}
	}

	return consts
}

// AddConst adds a constant to the file.
func (f *File) AddConst(c *Const) *File {
	f.Decls = append(f.Decls, c.ToDecl())

	return f
}

// AddVar adds a variable to the file.
func (f *File) AddVar(v *Var) *File {
	f.Decls = append(f.Decls, v.ToDecl())

	return f
}

// AddImport adds an import to the file.
func (f *File) AddImport(imp *Import) *File {
	if _, ok := f.toImport[imp.Path()]; ok {
		return f
	}

	f.Decls = append(f.Decls, imp.ToDecl())
	f.toImport[imp.Path()] = struct{}{}

	return f
}

// Print prints the file to the io.Writer.
func (f *File) Print(w io.Writer) error {
	return printer.Fprint(w, token.NewFileSet(), f.File)
}

// PrintFormatted prints the file to the io.Writer with imports.Process formatting.
func (f *File) PrintFormatted(w io.Writer) error {
	var buf bytes.Buffer
	if err := f.Print(&buf); err != nil {
		return err
	}

	if _, err := imports.Process("gen/tmp.go", buf.Bytes(), nil); err != nil {
		return err
	}

	return nil
}
