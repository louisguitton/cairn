package templates_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/louisguitton/cairn/internal/detector"
	"github.com/louisguitton/cairn/internal/templates"
)

var coreIDs = []string{"cairn-intake", "cairn-hill", "cairn-canvas", "cairn-handoff", "cairn-sync"}

func TestParseFrontmatterStageAndNext(t *testing.T) {
	meta := templates.ParseFrontmatter(`---
name: /cairn-hill
id: cairn-hill
category: Core
stage: reflect
next: /cairn-canvas (after the Hill playback)
description: test
---

body`)

	if meta.Stage != "reflect" {
		t.Errorf("Stage = %q", meta.Stage)
	}
	if !strings.HasPrefix(meta.Next, "/cairn-canvas") {
		t.Errorf("Next = %q", meta.Next)
	}
}

func TestEmbeddedCoreTemplates(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	core, err := mgr.ListCore()
	if err != nil {
		t.Fatal(err)
	}
	if len(core) != len(coreIDs) {
		t.Fatalf("core templates = %d, want %d", len(core), len(coreIDs))
	}

	byID := map[string]templates.TemplateMeta{}
	for _, tmpl := range core {
		byID[tmpl.ID] = tmpl
	}
	validStages := map[string]bool{"observe": true, "reflect": true, "make": true, "loop": true}
	for _, id := range coreIDs {
		tmpl, ok := byID[id]
		if !ok {
			t.Errorf("missing core template %s", id)
			continue
		}
		if !validStages[tmpl.Stage] {
			t.Errorf("%s stage = %q, invalid", id, tmpl.Stage)
		}
		if tmpl.Category != "Core" {
			t.Errorf("%s category = %q", id, tmpl.Category)
		}
		if id != "cairn-sync" && tmpl.Next == "" {
			t.Errorf("%s missing next: hint", id)
		}
	}
}

func TestBetaContainsReverse(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	beta, err := mgr.ListBeta()
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, tmpl := range beta {
		if tmpl.ID == "cairn-reverse" {
			found = true
			if tmpl.Category != "Beta" {
				t.Errorf("cairn-reverse category = %q", tmpl.Category)
			}
		}
	}
	if !found {
		t.Error("cairn-reverse not in beta templates")
	}
}

func TestFlatStrategyWritesCommands(t *testing.T) {
	dir := t.TempDir()
	mgr := templates.NewEmbeddedTemplateManager()

	results := templates.StrategyFor(detector.ClaudeCode, mgr).GenerateAll(dir, false)
	for _, r := range results {
		if !r.Success {
			t.Errorf("generate failed: %s", r.Message)
		}
	}

	for _, id := range coreIDs {
		path := filepath.Join(dir, ".claude/commands", id+".md")
		if _, err := os.Stat(path); err != nil {
			t.Errorf("missing %s", path)
		}
	}
}

func TestFlatStrategyRespectsForce(t *testing.T) {
	dir := t.TempDir()
	mgr := templates.NewEmbeddedTemplateManager()
	tmpl, err := mgr.GetByName("cairn-intake")
	if err != nil {
		t.Fatal(err)
	}

	strategy := templates.StrategyFor(detector.Cursor, mgr)
	first := strategy.GenerateOne(dir, tmpl, false)
	if !first[0].Success {
		t.Fatalf("first write failed: %s", first[0].Message)
	}
	second := strategy.GenerateOne(dir, tmpl, false)
	if second[0].Success {
		t.Error("second write without --force should fail")
	}
	third := strategy.GenerateOne(dir, tmpl, true)
	if !third[0].Success {
		t.Errorf("third write with force failed: %s", third[0].Message)
	}
}

func TestClaudeSkillBundle(t *testing.T) {
	dir := t.TempDir()
	mgr := templates.NewEmbeddedTemplateManager()

	results := templates.StrategyFor(detector.Claude, mgr).GenerateAll(dir, false)
	for _, r := range results {
		if !r.Success {
			t.Errorf("bundle write failed: %s (%s)", r.Message, r.FilePath)
		}
	}

	skillPath := filepath.Join(dir, "cairn-skill", "SKILL.md")
	data, err := os.ReadFile(skillPath)
	if err != nil {
		t.Fatalf("missing SKILL.md: %v", err)
	}
	content := string(data)
	if !strings.HasPrefix(content, "---\nname: cairn\n") {
		t.Error("SKILL.md missing frontmatter name")
	}
	for _, id := range coreIDs {
		if !strings.Contains(content, "references/"+id+".md") {
			t.Errorf("SKILL.md router missing %s", id)
		}
		refPath := filepath.Join(dir, "cairn-skill", "references", id+".md")
		refData, err := os.ReadFile(refPath)
		if err != nil {
			t.Errorf("missing reference %s", refPath)
			continue
		}
		if strings.HasPrefix(string(refData), "---") {
			t.Errorf("%s should have frontmatter stripped", refPath)
		}
	}
}

// Stage commands must never take outward-facing git actions on their own.
// A template that tells the agent to commit, push, or open a PR at the gate
// front-runs the human review that IS the gate.
func TestTemplatesForbidGitWritesAtTheGate(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	all, err := mgr.ListAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(all) == 0 {
		t.Fatal("no templates found")
	}

	for _, tmpl := range all {
		if !strings.Contains(tmpl.Content, "**Stop there — no git writes.**") {
			t.Errorf("%s: gate step must forbid git writes", tmpl.ID)
		}
		if !strings.Contains(tmpl.Content, "## Publishing (only on explicit request)") {
			t.Errorf("%s: missing the publishing-on-request section", tmpl.ID)
		}
		if strings.Contains(tmpl.Content, "per `gating.mode`") {
			t.Errorf("%s: gating.mode is descriptive, not an instruction to act", tmpl.ID)
		}
	}
}
