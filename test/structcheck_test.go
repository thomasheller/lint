package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestStructcheck(t *testing.T) {
	t.Parallel()
	source := `package test

type test struct {
	unused int
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 2}, Msg: "unused struct field github.com/thomasheller/mlint/regressiontests/.test.unused"},
	}
	ExpectIssues(t, "structcheck", source, expected)
}
