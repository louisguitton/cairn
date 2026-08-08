package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/louisguitton/cairn/internal/config"
)

func write(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestLoadMissingFileReturnsFallbackConfig(t *testing.T) {
	dir := t.TempDir()

	cfg, err := config.Load(dir)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	r := cfg.Resolve(config.Specs)
	if r.Bound {
		t.Errorf("expected unbound specs, got bound")
	}
	want := filepath.Join(dir, "design/canvas")
	if r.Path != want {
		t.Errorf("fallback path = %q, want %q", r.Path, want)
	}
}

func TestLoadSearchesUpward(t *testing.T) {
	root := t.TempDir()
	write(t, filepath.Join(root, config.FileName), "version: 1\nbindings:\n  hills: product/hills.md\n")
	nested := filepath.Join(root, "a", "b")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}

	cfg, err := config.Load(nested)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if cfg.Dir() != root {
		t.Errorf("Dir() = %q, want %q", cfg.Dir(), root)
	}
	if got := cfg.Resolve(config.Hills).Path; got != filepath.Join(root, "product/hills.md") {
		t.Errorf("hills path = %q", got)
	}
}

func TestResolveBetaBinding(t *testing.T) {
	proto := t.TempDir()
	docs := filepath.Join(proto, "..", filepath.Base(proto)+"-docs")
	if err := os.MkdirAll(docs, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(docs) })

	write(t, filepath.Join(docs, "product", "hills.md"), "# Hills")
	write(t, filepath.Join(proto, config.FileName), `version: 1
docs_repo:
  path: ../`+filepath.Base(docs)+`
proto_repo:
  path: .
bindings:
  hills: product/hills.md
  specs: product/specs
  capture: stream
gating:
  mode: pr
`)

	cfg, err := config.Load(proto)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}

	hills := cfg.Resolve(config.Hills)
	if !hills.Bound {
		t.Fatal("hills should be bound")
	}
	if !hills.ReadOnly {
		t.Error("bound hills must be read-only")
	}
	if !hills.Exists {
		t.Errorf("hills should exist at %s", hills.Path)
	}

	capture := cfg.Resolve(config.Capture)
	if capture.Exists {
		t.Error("capture path does not exist yet; Exists should be false")
	}
	if capture.ReadOnly {
		t.Error("capture must not be read-only")
	}

	if cfg.Gating.Mode != "pr" {
		t.Errorf("gating mode = %q, want pr", cfg.Gating.Mode)
	}
}

func TestResolveAllStableOrder(t *testing.T) {
	cfg, err := config.Load(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	rows := cfg.ResolveAll()
	if len(rows) != len(config.AllArtefactTypes) {
		t.Fatalf("rows = %d, want %d", len(rows), len(config.AllArtefactTypes))
	}
	for i, r := range rows {
		if r.Type != config.AllArtefactTypes[i] {
			t.Errorf("row %d = %s, want %s", i, r.Type, config.AllArtefactTypes[i])
		}
	}
}

func TestSaveRoundTrip(t *testing.T) {
	dir := t.TempDir()
	in := config.Config{
		Version:  1,
		DocsRepo: config.RepoRole{Path: "../docs"},
		Bindings: config.Bindings{Hills: "product/hills.md"},
		Gating:   config.Gating{Mode: "pr"},
	}

	path, err := config.Save(in, dir)
	if err != nil {
		t.Fatalf("Save: %v", err)
	}
	if filepath.Base(path) != config.FileName {
		t.Errorf("saved as %q", path)
	}

	out, err := config.Load(dir)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if out.DocsRepo.Path != "../docs" || out.Bindings.Hills != "product/hills.md" || out.Gating.Mode != "pr" {
		t.Errorf("round trip mismatch: %+v", out)
	}
}
