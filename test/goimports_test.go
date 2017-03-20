package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGoimports(t *testing.T) {
	source := `
package test
func test() {fmt.Println(nil)}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 1, Column: 0}, Msg: "file is not goimported"},
	}
	ExpectIssues(t, "goimports", source, expected)
}
