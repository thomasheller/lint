package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGoType(t *testing.T) {
	t.Parallel()
	source := `package test

func test() {
	var foo string
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 6}, Msg: "foo declared but not used"},
	}
	ExpectIssues(t, "gotype", source, expected)
}
