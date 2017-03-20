package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestMisSpell(t *testing.T) {
	t.Parallel()
	source := `package test
// The langauge is incorrect.
var a = "langauge"
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 2, Column: 7}, Msg: "\"langauge\" is a misspelling of \"language\""},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 9}, Msg: "\"langauge\" is a misspelling of \"language\""},
	}
	ExpectIssues(t, "misspell", source, expected)
}
