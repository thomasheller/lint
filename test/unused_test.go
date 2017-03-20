package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestUnused(t *testing.T) {
	t.Parallel()
	source := `package test

var v int = 10

func f() {
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 5}, Msg: "var v is unused"},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 5, Column: 6}, Msg: "func f is unused"},
	}
	ExpectIssues(t, "unused", source, expected)
}
