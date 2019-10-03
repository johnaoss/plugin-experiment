// Package testdata provides a couple interfaces to be used during the testing
// of the interpreted code.
package testdata

import (
	"fmt"
)

// Interface is an interface with a void function that doesn't have to do anything.
type Interface interface {
	Void()
}

// Struct is a base type that implements Interface
type Struct struct{}

// MakeStruct() returns a new struct
func MakeStruct() *Struct {
	return &Struct{}
}

// Void satisfies the interface 'Interface'
func (s *Struct) Void() {
	fmt.Println("Called 'Void' on Struct")
}
