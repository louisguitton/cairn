package cmd_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/louisguitton/cairn/internal/detector"
	"github.com/louisguitton/cairn/internal/templates"
)

// Beta templates are not installed by --all, but must be installable by name.
func TestBetaTemplateInstallableByName(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()

	tmpl, err := mgr.GetByName("cairn-reverse")
	if err != nil {
		t.Fatalf("beta template must resolve by name: %v", err)
	}
	if tmpl.Category != "Beta" {
		t.Errorf("category = %q, want Beta", tmpl.Category)
	}

	dir := t.TempDir()
	for _, r := range templates.StrategyFor(detector.ClaudeCode, mgr).GenerateOne(dir, tmpl, false) {
		if !r.Success {
			t.Fatalf("generate failed: %s", r.Message)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, ".claude/commands/cairn-reverse.md")); err != nil {
		t.Errorf("beta template not written: %v", err)
	}

	// --all stays core-only, so beta is opt-in.
	available, err := mgr.ListAvailable(detector.ClaudeCode)
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range available {
		if a.ID == "cairn-reverse" {
			t.Error("beta template must not be installed by --all")
		}
	}
}
