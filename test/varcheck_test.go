package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestVarcheck(t *testing.T) {
	t.Parallel()
	source := `package test

var v int
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 5}, Msg: "unused global variable v"},
	}
	ExpectIssues(t, "varcheck", source, expected)
}
