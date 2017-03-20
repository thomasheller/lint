package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestIneffassign(t *testing.T) {
	t.Parallel()
	source := `package test

func test() {
	a := 1
}`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 2}, Msg: "ineffectual assignment to a"},
	}
	ExpectIssues(t, "ineffassign", source, expected)
}
