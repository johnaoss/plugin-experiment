// Package foobar is an example of a type implementing an interface.
package foobar

import (
	"fmt"
)

// TT implements the testdata.Interface
type TT struct {
	A int
}

func New() *TT {
	return &TT{A: 40}
}

func (t *TT) Void() {
	fmt.Println("Called void on TT")
}
