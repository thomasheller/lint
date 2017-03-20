package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func TestDupl(t *testing.T) {
	t.Parallel()
	source := `package test

func findVendoredLinters() string {
	gopaths := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
	for _, home := range vendoredSearchPaths {
		for _, p := range gopaths {
			joined := append([]string{p, "src"}, home...)
			vendorRoot := filepath.Join(joined...)
			fmt.Println(vendorRoot)
			if _, err := os.Stat(vendorRoot); err == nil {
				return vendorRoot
			}
		}
	}
	return ""

}

func two() string {
	gopaths := strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator))
	for _, home := range vendoredSearchPaths {
		for _, p := range gopaths {
			joined := append([]string{p, "src"}, home...)
			vendorRoot := filepath.Join(joined...)
			fmt.Println(vendorRoot)
			if _, err := os.Stat(vendorRoot); err == nil {
				return vendorRoot
			}
		}
	}
	return ""

}
`

	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 19, Column: 0}, Msg: "duplicate of test.go:3-17"},
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 3, Column: 0}, Msg: "duplicate of test.go:19-33"},
	}
	ExpectIssues(t, "dupl", source, expected)
}
