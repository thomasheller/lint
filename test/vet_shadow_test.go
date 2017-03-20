package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestVetShadow(t *testing.T) {
	t.Parallel()
	source := `package test

func test(mystructs []*MyStruct) *MyStruct {
	var foo *MyStruct
	for _, mystruct := range mystructs {
		foo := mystruct
	}
	return foo
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 6, Column: 0}, Msg: "declaration of \"foo\" shadows declaration at test.go:4"},
	}
	ExpectIssues(t, "vetshadow", source, expected)
}
