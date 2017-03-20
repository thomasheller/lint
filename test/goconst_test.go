package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGoconst(t *testing.T) {
	t.Parallel()
	source := `package test
func a() {
	foo := "bar"
}
func b() {
	bar := "bar"
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 9}, Msg: `1 other occurrence(s) of "bar" found in: test.go:6:9`},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 6, Column: 9}, Msg: `1 other occurrence(s) of "bar" found in: test.go:3:9`},
	}
	ExpectIssues(t, "goconst", source, expected, "--min-occurrences", "2")
}
