package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGolint(t *testing.T) {
	t.Parallel()
	source := `
package test

type Foo int
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 6}, Msg: "exported type Foo should have comment or be unexported"},
	}
	ExpectIssues(t, "golint", source, expected)
}
