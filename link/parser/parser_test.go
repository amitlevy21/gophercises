package parser_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/amitlevy21/gophercises/link/parser"
)

// assert fails the test if the condition is false.
func assert(t *testing.T, condition bool, msg string) {
	if !condition {
		t.Logf(msg)
		t.FailNow()
	}
}

// ok fails the test if an err is not nil.
func ok(t *testing.T, err error) {
	if err != nil {
		t.Logf("unexpected error: %s\n\n", err.Error())
		t.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		t.Logf("\n\n\twant: %#v\n\n\tgot: %#v\n\n", exp, act)
		t.FailNow()
	}
}

var testData = []struct {
	file     string
	expected []parser.Link
}{
	{"ex1.html", []parser.Link{{Href: "/other-page", Text: "A link to another page"}}},
	{"ex2.html", []parser.Link{
		{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
		{Href: "https://github.com/gophercises", Text: "Gophercises is on <strong>Github</strong>!"}}},
}

func TestParse(t *testing.T) {
	for _, testCase := range testData {
		f, err := os.Open(filepath.Join("testdata", testCase.file))
		ok(t, err)

		equals(t, testCase.expected, parser.Parse(f))
	}
}
