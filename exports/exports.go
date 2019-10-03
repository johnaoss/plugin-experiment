package exports

import "reflect"

// Symbols represents the entirety of the exported types of this repo.
var Symbols = map[string]map[string]reflect.Value{}

// TestDataSymbols is the wrapper around the testdata interface. This only contains
// The symbols needed for wrapping the testdata interface.
var TestDataSymbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/johnaoss/plugin-experiment/exports"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}

	TestDataSymbols["github.com/johnaoss/plugin-experiment/testdata"] = Symbols["github.com/johnaoss/plugin-experiment/testdata"]
}

//go:generate goexports github.com/johnaoss/plugin-experiment/testdata
//go:generate goexports github.com/johnaoss/plugin-experiment/testdata/foobar
