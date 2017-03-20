package regressiontests

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/kisielk/gotool"
	"go/types"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa/ssautil"

	"github.com/stretchr/testify/assert"

	"github.com/thomasheller/lint"

	// import linters - requires a patched version of each linter
	// that implements Checker:
	golint "github.com/golang/lint"
	"github.com/kisielk/errcheck"
	"github.com/mdempsky/unconvert"
	"github.com/mvdan/interfacer/cmd/interfacer"
	"github.com/opennota/check/aligncheck"
	"github.com/opennota/check/structcheck"
	"github.com/opennota/check/varcheck"
	"github.com/tsenart/deadcode"
	"honnef.co/go/tools/cmd/gosimple"
	"honnef.co/go/tools/cmd/staticcheck"
	"honnef.co/go/tools/cmd/unused"
)

func ExpectIssues(t *testing.T, linter string, source string, expected lint.Issues, extraFlags ...string) {
	dir, err := ioutil.TempDir(".", "tmp-")
	if !assert.NoError(t, err) {
		return
	}
	defer os.RemoveAll(dir)

	f := filepath.Join(dir, "test.go")

	abs, err := filepath.Abs(f)
	if !assert.NoError(t, err) {
		return
	}

	w, err := os.Create(f)

	if !assert.NoError(t, err) {
		return
	}
	defer os.Remove(w.Name())
	_, err = w.WriteString(source)
	_ = w.Close()
	if !assert.NoError(t, err) {
		return
	}

	paths := gotool.ImportPaths([]string{w.Name()})
	var conf loader.Config
	if _, err := conf.FromArgs(paths, true); err != nil {
		t.Error(err)
	}
	lprog, err := conf.Load()
	if err != nil {
		t.Error(err)
	}
	wantPkg := make(map[*types.Package]bool)
	for _, info := range lprog.InitialPackages() {
		wantPkg[info.Pkg] = true
	}
	prog := ssautil.CreateProgram(lprog, 0)
	prog.Build()

	var actual lint.Issues

	switch linter {
	case "aligncheck":
		actual, err = aligncheck.Check(lprog, prog)
	case "deadcode":
		actual, err = deadcode.Check(lprog, prog)
	case "errcheck":
		actual, err = errcheck.Check(lprog, prog)
	case "golint":
		actual, err = golint.Check(lprog, prog)
	case "gosimple":
		actual, err = gosimple.Check(lprog, prog)
	case "interfacer":
		actual, err = interfacer.Check(lprog, prog)
	case "staticcheck":
		actual, err = staticcheck.Check(lprog, prog)
	case "structcheck":
		actual, err = structcheck.Check(lprog, prog)
	case "unconvert":
		actual, err = unconvert.Check(lprog, prog)
	case "unused":
		actual, err = unused.Check(lprog, prog)
	case "varcheck":
		actual, err = varcheck.Check(lprog, prog)
	default:
		t.Fatalf("linter %s not supported", linter)
	}

	actualForCmp := lint.Issues{}
	for _, issue := range actual {
		issue.Position.Filename = "test.go"
		issue.Msg = strings.Replace(issue.Msg, abs, "test.go", -1)
		actualForCmp = append(actualForCmp, issue)
	}

	sort.Sort(expected)
	sort.Sort(actualForCmp)

	assert.Equal(t, expected, actualForCmp)
}
