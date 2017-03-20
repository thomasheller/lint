package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGofmt(t *testing.T) {
	t.Parallel()
	source := `
package test
func test() { if nil {} }
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 1, Column: 0}, Msg: "file is not gofmted with -s"},
	}
	ExpectIssues(t, "gofmt", source, expected)
}
