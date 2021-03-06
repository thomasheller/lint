// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package lint defines common interfaces for Go code checkers.
package lint

import (
	"go/token"

	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
)

// A Checker points out issues in a program.
type Checker interface {
	Check(*loader.Program, *ssa.Program) ([]Issue, error)
}

// Issue represents an issue somewhere in a source code file.
type Issue struct {
	Pos token.Pos
	Msg string
}
