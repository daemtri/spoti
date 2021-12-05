package apispot

import (
	"embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"path/filepath"
	"strings"
)

func Parser(specific embed.FS) {
	parserFiles(specific)
}

func parserFiles(specific embed.FS) error {
	fileset := token.NewFileSet()
	pkgs, first := parseDir(fileset, specific, ".", nil, parser.ParseComments|parser.AllErrors)
	if first != nil {
		return fmt.Errorf("parser dir error: %w", first)
	}

	v := &astVisitor{
		Imports: map[string]string{},
	}

	for _, aPackage := range pkgs {
		parsePackage(aPackage, v)
	}

	return nil
}

func checkImports(aFile *ast.File) error {
	for _, ip := range aFile.Imports {
		path := ip.Path.Value
		if path != "context" && path != "errors" && path != "embed" && path != "json" {
			return fmt.Errorf("%s 包含不允许的import %s", aFile.Name, path)
		}
	}
	return nil
}

func parsePackage(aPackage *ast.Package, v *astVisitor) error {
	for key, fileEntry := range aPackage.Files {
		if err := checkImports(fileEntry); err != nil {
			return err
		}
		v.CurrentFilename = key
		ast.Walk(v, fileEntry)
	}
	return nil
}

func parseDir(fset *token.FileSet, dir embed.FS, path string, filter func(fs.FileInfo) bool, mode parser.Mode) (pkgs map[string]*ast.Package, first error) {
	list, err := dir.ReadDir(path)
	if err != nil {
		return nil, err
	}

	pkgs = make(map[string]*ast.Package)
	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		if filter != nil {
			info, err := d.Info()
			if err != nil {
				return nil, err
			}
			if !filter(info) {
				continue
			}
		}
		filename := filepath.Join(path, d.Name())
		fileBytes, err := dir.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		if src, err := parser.ParseFile(fset, filename, fileBytes, mode); err == nil {
			name := src.Name.Name
			pkg, found := pkgs[name]
			if !found {
				pkg = &ast.Package{
					Name:  name,
					Files: make(map[string]*ast.File),
				}
				pkgs[name] = pkg
			}
			pkg.Files[filename] = src
		} else if first == nil {
			first = err
		}
	}

	return
}

type astVisitor struct {
	CurrentFilename string
	PackageName     string
	Filename        string
	Imports         map[string]string
	Arguments       []Struct
	Replies         []Struct
	Service         Interface
	Enums           []Enum
}

func (av *astVisitor) Visit(node ast.Node) (w ast.Visitor) {
	return nil
}
