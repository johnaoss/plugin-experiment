package exports_test

import (
	"reflect"
	"testing"

	"github.com/johnaoss/plugin-experiment/exports"
)

// TestEmptyExports is used as a sanity check to verify exports have been
// generated successfully.
func TestEmptyExports(t *testing.T) {
	if len(exports.Symbols) == 0 {
		t.Error("Symbols should contain _some_ elements but has size 0.")
	}

	if len(exports.TestDataSymbols) == 0 {
		t.Error("TestDataSymbols should contain _some_ elements but has size 0.")
	}

	if t.Failed() {
		t.Error("Ensure that `go generate` was ran for this package.")
		return
	}

	if testing.Verbose() {
		printSymbols(t, exports.Symbols)
	}
}

// printSymbols is a helper function to print what's exactly in the reflect.Value map.
// <self> is shown if the map is containing a shallow copy of symbols.
func printSymbols(t *testing.T, syms map[string]map[string]reflect.Value) {
	t.Helper()
	for k, v := range syms {
		t.Logf("Package: %s", k)
		for i, j := range v {
			if k == "github.com/johnaoss/plugin-experiment/exports" && i == "Symbols" {
				t.Logf("\tKey: <%s> Value: %v", i, "<self>")
				continue
			}
			t.Logf("\tKey: <%s> Value: %v", i, j)
		}
	}
}
