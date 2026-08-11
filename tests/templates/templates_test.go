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
		if !strings.Contains(tmpl.Content, "**Stop there. No git writes.**") {
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

// Stages must share one work identity. Intake names the work; everything
// downstream inherits that name from the artefact header. Re-deriving a slug
// per stage is what scattered the first real run across three identities.
func TestTemplatesShareOneWorkIdentity(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	all, err := mgr.ListAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, tmpl := range all {
		establishes := strings.Contains(tmpl.Content, "Establish the work identity") ||
			strings.Contains(tmpl.Content, "identity the way `/cairn-intake` does")
		inherits := strings.Contains(tmpl.Content, "work identity from the input artefact's `Work` header") ||
			strings.Contains(tmpl.Content, "taken from the input artefact's `Work` header")

		if !establishes && !inherits {
			t.Errorf("%s: must either establish the work identity or inherit it", tmpl.ID)
		}
	}
}

// The feedback that drove these rules: artefacts full of codes, nine-column
// decision tables and rhetorical prose did not survive a real review meeting
// with non-native English speakers. Prompt style is contagious, so the
// templates must obey the rules they impose.
func TestTemplatesObeyTheirOwnWritingRules(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	all, err := mgr.ListAll()
	if err != nil {
		t.Fatal(err)
	}

	for _, tmpl := range all {
		// An em dash in the instructions teaches the model to use em dashes.
		if strings.Contains(tmpl.Content, "\u2014") {
			t.Errorf("%s: contains an em dash, which the writing rules ban", tmpl.ID)
		}
		// Confidence must be words, never letter codes.
		for _, code := range []string{"K/A/?", "(K)", "(A)", "· K", "· A"} {
			if strings.Contains(tmpl.Content, code) {
				t.Errorf("%s: uses the %q confidence code instead of a word", tmpl.ID, code)
			}
		}
		if !strings.Contains(tmpl.Content, "## Writing rules") {
			t.Errorf("%s: missing the writing-rules section", tmpl.ID)
		}
		if !strings.Contains(tmpl.Content, "Never refer to anything by an identifier alone") {
			t.Errorf("%s: missing the ban on identifier-only cross-references", tmpl.ID)
		}
		if !strings.Contains(tmpl.Content, "Simplify the sentence, never the claim") {
			t.Errorf("%s: missing the rule protecting claims from readability edits", tmpl.ID)
		}
		if !strings.Contains(tmpl.Content, "Protect the uncomfortable findings") {
			t.Errorf("%s: missing the guardrail protecting honest findings", tmpl.ID)
		}
	}
}

// Decisions are read aloud one at a time, so they are blocks, not table rows.
func TestDecisionsUseBlockFormat(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	for _, id := range []string{"cairn-canvas", "cairn-sync", "cairn-reverse"} {
		tmpl, err := mgr.GetByName(id)
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		for _, line := range []string{"**Decided.**", "**Why.**", "**Rejected.**", "**What it costs.**", "**Who owns it.**"} {
			if !strings.Contains(tmpl.Content, line) {
				t.Errorf("%s: decision block missing the %s line", id, line)
			}
		}
		if strings.Contains(tmpl.Content, "| ID  | Date") {
			t.Errorf("%s: still carries the old decision table", id)
		}
	}
}

// Twelve unranked questions is a backlog, not a list.
func TestOpenQuestionsAreTriaged(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()
	tmpl, err := mgr.GetByName("cairn-canvas")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"### Critical", "can invalidate the work or block downstream work", "more than about eight open questions"} {
		if !strings.Contains(tmpl.Content, want) {
			t.Errorf("cairn-canvas: open questions missing %q", want)
		}
	}
}

// The folder listing is the review agenda, so artefacts are numbered in the
// order they should be read. The hill is always reviewed before the canvas.
func TestArtefactFilenamesAreNumberedByReviewOrder(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()

	writes := map[string][]string{
		"cairn-intake":  {"`<home>/1-intake.md`"},
		"cairn-hill":    {"`<home>/2-hill.md`"},
		"cairn-canvas":  {"`<home>/3-canvas.md`"},
		"cairn-handoff": {"`<home>/4-brief.md`", "`<home>/5-stories.md`"},
		"cairn-sync":    {"`<home>/3-canvas.md`"},
		"cairn-reverse": {"`<home>/2-hill.md`", "`<home>/3-canvas.md`"},
	}

	for id, paths := range writes {
		tmpl, err := mgr.GetByName(id)
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		for _, want := range paths {
			if !strings.Contains(tmpl.Content, want) {
				t.Errorf("%s: expected to write %s", id, want)
			}
		}
		// An unnumbered path would sort wrong in the folder listing.
		for _, bare := range []string{"`<home>/intake.md`", "`<home>/hill.md`", "`<home>/canvas.md`", "`<home>/brief.md`", "`<home>/stories.md`"} {
			if strings.Contains(tmpl.Content, bare) {
				t.Errorf("%s: still writes the unnumbered path %s", id, bare)
			}
		}
	}
}

// A title has to stand on its own: the folder listing and any later reference
// show the title, not the folder path. Emergent good titles are not enough,
// because nothing stops them regressing.
func TestTitlesMustBeSelfContainedPlainWords(t *testing.T) {
	mgr := templates.NewEmbeddedTemplateManager()

	// Every template that writes an artefact with a `# X: {title}` heading.
	for _, id := range []string{"cairn-intake", "cairn-hill", "cairn-canvas", "cairn-handoff", "cairn-reverse"} {
		tmpl, err := mgr.GetByName(id)
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		if !strings.Contains(tmpl.Content, "The title is one line of plain words that stands on its own") {
			t.Errorf("%s: missing the self-contained title rule", id)
		}
		if !strings.Contains(tmpl.Content, "State the outcome for the user, not the mechanism") {
			t.Errorf("%s: title rule must demand the outcome over the mechanism", id)
		}
	}

	// The slug is derived from the title, so the two cannot disagree.
	intake, err := mgr.GetByName("cairn-intake")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(intake.Content, "kebab-case words lifted from the title") {
		t.Error("cairn-intake: the slug must be derived from the title")
	}
}
