// Package trimimports is a libary to remove unused Go imports from (generated) source code.
package trimimports

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// Trim removes any unused imports from the given file.
func Trim(fset *token.FileSet, f *ast.File) {
	for _, importBlock := range astutil.Imports(fset, f) {
		for _, s := range importBlock {
			path := strings.Trim(s.Path.Value, `"`)
			if !astutil.UsesImport(f, path) {
				if s.Name != nil {
					astutil.DeleteNamedImport(fset, f, s.Name.Name, path)
				} else {
					astutil.DeleteImport(fset, f, path)
				}
			}
		}
	}
}

// TrimAndReformat parses, Trim()s and (re)formats the given Go code.
func TrimAndReformat(src []byte) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "generated.go", src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	Trim(fset, f)

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
