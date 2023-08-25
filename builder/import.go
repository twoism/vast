package builder

import (
	"go/ast"
	"go/token"
	"strings"
)

type Import struct {
	*ast.ImportSpec

	baseName string
	path     string
	alias    string
}

type ImportOpt func(*Import)

func ImportAliasOpt(alias string) ImportOpt {
	return func(i *Import) {
		i.alias = alias
		i.Name = ast.NewIdent(alias)
	}
}

// ParsePackage parses the package from the given path.
func ParsePackage(path string) (pkg, base string, is3Pkg bool) {
	if strings.Contains(path, "/") {
		parts := strings.Split(path, "/")
		base = parts[len(parts)-1]
		pkg = path
	} else {
		pkg = path
		base = pkg
	}

	return
}

// NewImport creates a new import with the given path.
func NewImport(path string, opts ...ImportOpt) *Import {
	pkg, base, _ := ParsePackage(path)

	imp := &Import{
		path:     pkg,
		baseName: base,
		ImportSpec: &ast.ImportSpec{
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"" + path + "\"",
			},
		},
	}

	for _, opt := range opts {
		opt(imp)
	}

	return imp
}

// BaseName returns the base name of the import.
func (i *Import) BaseName() string {
	if i.baseName == "" {
		i.baseName = i.Path()[1 : len(i.Path())-1]
	}

	return i.baseName
}

// Path returns the path of the import.
func (i *Import) Path() string {
	return i.path
}

// Alias returns the alias of the import.
func (i *Import) Alias() string {
	return i.alias
}

// HasAlias returns true if the import has an alias.
func (i *Import) HasAlias() bool {
	return i.alias != ""
}

// ToDecl returns the import as a declaration.
func (i *Import) ToDecl() ast.Decl {
	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{i.ImportSpec},
	}
}
