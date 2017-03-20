package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestDeadcode(t *testing.T) {
	t.Parallel()
	source := `package test

func test() {
	return
	println("hello")
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 1}, Msg: "test is unused"},
	}
	ExpectIssues(t, "deadcode", source, expected)
}
