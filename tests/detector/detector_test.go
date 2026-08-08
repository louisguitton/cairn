package detector_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/louisguitton/cairn/internal/detector"
)

func TestDetectClaudeCode(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".claude"), 0o755); err != nil {
		t.Fatal(err)
	}

	d := detector.NewDefaultDetector()
	res := d.Detect(dir)

	if !res.IsValid || res.ToolType != detector.ClaudeCode {
		t.Fatalf("got %+v, want ClaudeCode", res)
	}
	if res.ConfigPath != filepath.Join(dir, ".claude/commands") {
		t.Errorf("config path = %q", res.ConfigPath)
	}
}

func TestDetectCursor(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, ".cursorrules"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}

	res := detector.NewDefaultDetector().Detect(dir)
	if !res.IsValid || res.ToolType != detector.Cursor {
		t.Fatalf("got %+v, want Cursor", res)
	}
}

func TestClaudeIsNeverAutoDetected(t *testing.T) {
	if sigs := detector.Claude.GetSignatureFiles(); sigs != nil {
		t.Errorf("Claude signatures = %v, want nil", sigs)
	}

	res := detector.NewDefaultDetector().Detect(t.TempDir())
	if res.IsValid {
		t.Errorf("empty dir should not detect anything, got %+v", res)
	}
}

func TestConfigDirs(t *testing.T) {
	cases := map[detector.AIToolType]string{
		detector.ClaudeCode: ".claude/commands",
		detector.Cursor:     ".cursor/commands",
		detector.Claude:     "cairn-skill",
		detector.Unknown:    "",
	}
	for tool, want := range cases {
		if got := tool.GetConfigDir(); got != want {
			t.Errorf("%s config dir = %q, want %q", tool, got, want)
		}
	}
}
