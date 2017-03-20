package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestVet(t *testing.T) {
	t.Parallel()
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 7, Column: 0}, Msg: "missing argument for Printf(\"%d\"): format reads arg 1, have only 0 args"},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 7, Column: 0}, Msg: "unreachable code"},
	}
	ExpectIssues(t, "vet", `package main

import "fmt"

func main() {
	return
	fmt.Printf("%d")
}
`, expected)
}
