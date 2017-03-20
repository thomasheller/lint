package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestUnconvert(t *testing.T) {
	t.Parallel()
	source := `package test

func test() {
	var a int64
	b := int64(a)
	println(b)
}`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 5, Column: 12}, Msg: "unnecessary conversion"},
	}
	ExpectIssues(t, "unconvert", source, expected)
}
