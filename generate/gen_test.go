package generate_test

import (
	"github.com/johnaoss/plugin-experiment/exports"
	"go/build"
	"reflect"
	"testing"

	"github.com/containous/yaegi/interp"
	"github.com/containous/yaegi/stdlib"

	"github.com/johnaoss/plugin-experiment/testdata"
	fbar "github.com/johnaoss/plugin-experiment/testdata/foobar"
)

var (
	// iType is the reflect.Type of a testdata.Interface
	iType = reflect.TypeOf((*testdata.Interface)(nil)).Elem()
)

// TestParseFailsNoSymbols is used to verify that the file panics if interface
// definitions are not defined at compile time. This test is expected to panic
// in it's current state.
//
// Note: this seems like this results due to the interface wrapper not being
// properly generated or interpreted by Yaegi.
func TestFoobarCanParse(t *testing.T) {
	parser := interp.New(interp.Options{GoPath: build.Default.GOPATH})
	// the Go stdlib symbols
	parser.Use(stdlib.Symbols)
	// the exported symbols of the testdata.interface
	parser.Use(exports.TestDataSymbols)

	// We expect this to panic
	defer func() {
		r := recover()
		if testing.Verbose() {
			t.Log("Panic Result:", r)
		}
		if r == nil {
			t.Fatal("Didn't need to recover")
		}
	}()

	// Assert we can import the testdata interface without fail.
	_, err := parser.Eval(`import "github.com/johnaoss/plugin-experiment/testdata"`)
	if err != nil {
		t.Fatalf("Failed to import testdata: %v", err)
	}

	// Assert we can import the testdata from exports.
	_, err = parser.Eval(`import "github.com/johnaoss/plugin-experiment/testdata/foobar"`)
	if err != nil {
		t.Fatal("Couldn't import foobar", err.Error())
	}

	// Now verify we can get the New() function.
	fn, err := parser.Eval("foobar.New")
	if err != nil {
		t.Fatalf("Failed to access New, given err: %v", err)
	}

	// If it can't interface, will panic.
	if !fn.CanInterface() {
		t.Fatalf("Function can't interface, aborting.")
	}

	iresult := fn.Interface().(func() *struct{ A int })()
	if testing.Verbose() {
		t.Log("TypeOf Struct:", reflect.TypeOf(iresult).String())
		t.Log("Struct # Methods:", reflect.ValueOf(iresult).NumMethod())
		t.Log("Struct implements testdata.Interface:", reflect.ValueOf(iresult).Type().Implements(iType))
	}

	// This line panics, may be able to generate a function that returns the type based on the struct fields.
	// But that's rather hacky...
	result := fn.Interface().(func() *fbar.TT)()

	if testing.Verbose() {
		t.Log("TypeOf Foobar:", reflect.TypeOf(result).String())
		t.Log("Foobar # Methods:", reflect.ValueOf(result).NumMethod())
		t.Log("Foobar implements testdata.Interface:", reflect.ValueOf(result).Type().Implements(iType))
	}

}

// TestInterfaceSatisfied checks if we can satisfy the interface after getting the
// interface wrappers at compile time.
func TestInterfaceSatisfied(t *testing.T) {
	parser := interp.New(interp.Options{GoPath: build.Default.GOPATH})
	// the Go stdlib symbols
	parser.Use(stdlib.Symbols)
	// the exported symbols of the testdata.interface
	parser.Use(exports.Symbols)

	// Assert we can import the testdata interface without fail.
	_, err := parser.Eval(`import "github.com/johnaoss/plugin-experiment/testdata"`)
	if err != nil {
		t.Fatalf("Failed to import testdata: %v", err)
	}

	// Now we wish to get the func to get a new struct
	sfunc, err := parser.Eval(`testdata.MakeStruct`)
	if err != nil {
		t.Fatalf("Failed to acquire MakeStruct function: %v", err)
	}

	result := sfunc.Interface().(func() *testdata.Struct)()

	result.Void()

}

// TestInterpreterHappyPath determines if we can properly use the interpreter.
// Essentially just a sanity check.
func TestInterpreterHappyPath(t *testing.T) {
	parser := interp.New(interp.Options{GoPath: build.Default.GOPATH})
	// the Go stdlib symbols
	parser.Use(stdlib.Symbols)

	// Sanity check if we can output "Hello World!""
	_, err := parser.Eval(`import "fmt"`)
	if err != nil {
		t.Fatal("Couldn't import fmt", err.Error())
	}
	_, err = parser.Eval(`fmt.Println("Hello world!")`)
	if err != nil {
		t.Fatal("Couldn't hello world", err.Error())
	}
}

// TestFoobarInterface just double checks that the foobar type implements
// the testdata.Inferface.
func TestFoobarInterface(t *testing.T) {
	fb := fbar.New()
	iType := reflect.TypeOf((*testdata.Interface)(nil)).Elem()

	if !reflect.TypeOf(fb).Implements(iType) {
		t.Error("Fbar does not implement testdata.Interface")
	}
}
