package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestStaticCheck(t *testing.T) {
	t.Parallel()
	source := `package test

import "regexp"

var v = regexp.MustCompile("*")
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 5, Column: 28}, Msg: "error parsing regexp: missing argument to repetition operator: `*`"},
	}
	ExpectIssues(t, "staticcheck", source, expected)
}
