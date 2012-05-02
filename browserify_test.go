package browserify_test

import (
	"github.com/nshah/go.browserify"
	"go/build"
	"strings"
	"testing"
)

const example = "github.com/nshah/go.browserify/example"

func exampleDir(t *testing.T) string {
	pkg, err := build.Import(example, "", build.FindOnly)
	if err != nil {
		t.Fatalf("Failed to find example npm module: %s", err)
	}
	return pkg.Dir
}

func TestContents(t *testing.T) {
	s := browserify.Script{
		Dir:   exampleDir(t),
		Entry: "lib/example.js",
	}
	b, err := s.Content()
	if err != nil {
		t.Fatalf("Error getting content: %s", err)
	}
	content := string(b)
	if !strings.Contains(content, "dotaccess") {
		t.Fatalf("Was expecting dotaccess but did not find it:\n%s", content)
	}
	if !strings.Contains(content, "42 + ans") {
		t.Fatalf("Was expecting example content but did not find it:\n%s", content)
	}
}
