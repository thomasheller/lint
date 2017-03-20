package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestAlignCheck(t *testing.T) {
	t.Parallel()
	source := `package test

type unaligned struct {
	a uint16
	b uint64
	c uint16

}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 19, Line: 3, Column: 6}, Msg: "test: test.go:3:6: struct unaligned could have size 16 (currently 24)"},
	}
	ExpectIssues(t, "aligncheck", source, expected)
}
