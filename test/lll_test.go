package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestLLL(t *testing.T) {
	t.Parallel()
	source := `package test
// This is a really long line full of text that is uninteresting in the extreme. Also we're just trying to make it here.
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 2, Column: 0}, Msg: "line is 120 characters"},
	}
	ExpectIssues(t, "lll", source, expected)
}
