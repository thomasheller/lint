package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestErrcheck(t *testing.T) {
	t.Parallel()
	source := `package moo

func f() error { return nil}
func test() { f() }
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 16}, Msg: "error return value not checked (func test() { f() })"},
	}
	ExpectIssues(t, "errcheck", source, expected)
}
