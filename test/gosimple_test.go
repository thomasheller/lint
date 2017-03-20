package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestGoSimple(t *testing.T) {
	t.Parallel()
	source := `package test

func a(ok bool, ch chan bool) {
	select {
	case <- ch:
	}

	for {
		select {
		case <- ch:
		}
	}

	if ok == true {
	}
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 8, Column: 2}, Msg: "should use for range instead of for { select {} }"},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 14, Column: 5}, Msg: "should omit comparison to bool constant, can be simplified to ok"},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 4, Column: 2}, Msg: "should use a simple channel send/receive instead of select with a single case"},
	}
	ExpectIssues(t, "gosimple", source, expected)
}
