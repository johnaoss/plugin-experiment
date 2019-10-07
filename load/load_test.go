package load_test

import (
	"bytes"
	"go/ast"
	"go/printer"
	"path/filepath"
	"testing"

	"github.com/johnaoss/plugin-experiment/load"
)

// TestLoad tests the loading of a given Go file, and the rewriting of its
// package name.
func TestLoad(t *testing.T) {
	// This loads a package or an individual file in said package.
	path, err := filepath.Abs("./_test/foo/")
	if err != nil {
		t.Fatalf("failed to get filepath: %v", err)
	}

	pkgs, err := load.LoadFile(path + "/foo.go")
	if err != nil {
		t.Fatalf("Failed to load file: %v", err)
	}

	// Rewrite the package declaration to be "bar"
	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			file.Name = ast.NewIdent("bar")

			var buf bytes.Buffer
			printer.Fprint(&buf, pkg.Fset, file)
			t.Log(buf.String())
		}
	}
}
