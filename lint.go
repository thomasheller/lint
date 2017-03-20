// Copyright (c) 2017, Daniel Martí <mvdan@mvdan.cc>
// See LICENSE for licensing information

// Package lint defines common interfaces for Go code checkers.
package lint

import (
	"fmt"
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
	Position token.Position
	Msg      string
}

func (i *Issue) String() string {
	return fmt.Sprintf("%s:%d:%d: %s",
		i.Position.Filename,
		i.Position.Line,
		i.Position.Column,
		i.Msg,
	)
}

type Issues []Issue

func (e Issues) Len() int           { return len(e) }
func (e Issues) Less(a, b int) bool { return e[a].String() < e[b].String() }
func (e Issues) Swap(a, b int)      { e[a], e[b] = e[b], e[a] }
