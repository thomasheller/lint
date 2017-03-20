package regressiontests

import (
	"github.com/thomasheller/lint"
	"go/token"
	"testing"
)

func IgnoreTestSafesql(t *testing.T) {
	t.Parallel()
	source := `package test

import (
	"database/sql"
	"log"
	"strconv"
)

func main() {
	getUser(42)
}

func getUser(userID int64) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM users WHERE id=" + strconv.FormatInt(userID, 10))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
}
`
	expected := lint.Issues{
		{Position: token.Position{Filename: "test.go", Offset: 0, Line: 20, Column: 23}, Msg: `potentially unsafe SQL statement`},
	}
	ExpectIssues(t, "safesql", source, expected)
}
