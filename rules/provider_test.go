package rules

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test_RulesAreDocumented ensures every rule has a unique name, a
// documentation page and an entry in the rules index.
func Test_RulesAreDocumented(t *testing.T) {
	index, err := os.ReadFile(filepath.Join("..", "docs", "rules", "README.md"))
	if err != nil {
		t.Fatalf("failed to read the rules index: %s", err)
	}

	seen := map[string]bool{}
	for _, rule := range Rules {
		name := rule.Name()

		if seen[name] {
			t.Errorf("rule %s is registered twice", name)
		}
		seen[name] = true

		if !strings.HasPrefix(name, "scaleway_") {
			t.Errorf("rule %s must be prefixed with scaleway_", name)
		}

		if _, err := os.Stat(filepath.Join("..", "docs", "rules", name+".md")); err != nil {
			t.Errorf("rule %s has no documentation page: %s", name, err)
		}

		if !strings.Contains(string(index), "["+name+"]("+name+".md)") {
			t.Errorf("rule %s is missing from docs/rules/README.md", name)
		}

		if !strings.HasSuffix(rule.Link(), "/docs/rules/"+name+".md") {
			t.Errorf("rule %s has an unexpected link: %s", name, rule.Link())
		}
	}
}
