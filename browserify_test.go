package browserify_test

import (
	"github.com/nshah/go.browserify"
	"testing"
)

func TestContents(t *testing.T) {
	s := browserify.Script{
		Dir:   "example",
		Entry: "lib/example.js",
	}
	content, err := s.Content()
	if err != nil {
		t.Fatalf("Error getting content %s", err)
	}
	t.Fatal(string(content))
}
