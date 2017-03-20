package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestInterfacer(t *testing.T) {
	t.Parallel()
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 5, Column: 8}, Msg: "r can be io.Closer"},
	}
	ExpectIssues(t, "interfacer", `package main

import "os"

func f(r *os.File) {
	r.Close()
}

func main() {
}
`, expected)
}
