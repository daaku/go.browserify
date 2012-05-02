package browserify_test

import (
	"github.com/nshah/go.browserify"
	"strings"
	"testing"
)

func TestContents(t *testing.T) {
	s := browserify.Script{
		Dir:   "example",
		Entry: "lib/example.js",
	}
	b, err := s.Content()
	if err != nil {
		t.Fatalf("Error getting content %s", err)
	}
	content := string(b)
	if !strings.Contains(content, "dotaccess") {
		t.Fatalf("Was expecting dotaccess but did not find it:\n%s", content)
	}
	if !strings.Contains(content, "42 + ans") {
		t.Fatalf("Was expecting example content but did not find it:\n%s", content)
	}
}
