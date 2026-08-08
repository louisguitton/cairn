package templates

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/louisguitton/cairn/internal"
	"github.com/louisguitton/cairn/internal/detector"
)

// ClaudeSkillBundleStrategy emits a single uploadable skill bundle for
// claude.ai / Claude Desktop, which has no per-repo commands directory:
//
//	cairn-skill/
//	  SKILL.md              — method overview + stage router
//	  references/<id>.md    — one file per command template
//
// The bundle is installed manually (upload in claude.ai skill settings);
// re-running `cairn generate --tool claude --force` refreshes it.
type ClaudeSkillBundleStrategy struct {
	manager *EmbeddedTemplateManager
}

func init() {
	RegisterStrategy(detector.Claude, func(mgr *EmbeddedTemplateManager) GenerationStrategy {
		return &ClaudeSkillBundleStrategy{manager: mgr}
	})
}

func (s *ClaudeSkillBundleStrategy) GenerateAll(workingDir string, force bool) []GenerateResult {
	tmpls, err := s.manager.ListAvailable(detector.Claude)
	if err != nil {
		return []GenerateResult{{
			Success: false,
			Message: "failed to list templates: " + err.Error(),
			Error:   err,
		}}
	}

	results := make([]GenerateResult, 0, len(tmpls)+1)
	results = append(results, s.writeRouter(workingDir, tmpls, force))
	for _, tmpl := range tmpls {
		results = append(results, s.GenerateOne(workingDir, tmpl, force)...)
	}
	return results
}

func (s *ClaudeSkillBundleStrategy) GenerateOne(workingDir string, tmpl TemplateMeta, force bool) []GenerateResult {
	refDir := filepath.Join(workingDir, detector.Claude.GetConfigDir(), "references")
	if err := os.MkdirAll(refDir, 0o755); err != nil {
		return []GenerateResult{{
			Success:  false,
			FilePath: refDir,
			Message:  "failed to create references directory: " + err.Error(),
			Error:    err,
		}}
	}

	refPath := filepath.Join(refDir, tmpl.ID+".md")
	return []GenerateResult{writeBundleFile(refPath, []byte(stripFrontmatter(tmpl.Content)), force)}
}

func (s *ClaudeSkillBundleStrategy) writeRouter(workingDir string, tmpls []TemplateMeta, force bool) GenerateResult {
	skillDir := filepath.Join(workingDir, detector.Claude.GetConfigDir())
	if err := os.MkdirAll(skillDir, 0o755); err != nil {
		return GenerateResult{
			Success:  false,
			FilePath: skillDir,
			Message:  "failed to create skill directory: " + err.Error(),
			Error:    err,
		}
	}

	var b strings.Builder
	b.WriteString("---\n")
	b.WriteString("name: cairn\n")
	b.WriteString(`description: "Prompt-driven Product Design Thinking on the IBM EDT Loop (Observe → Reflect → Make). Use when the user wants to run a cairn stage: intake a stakeholder request, draft a hill, assemble a design canvas, hand off to prototyping/stories, or sync reality (diffs, playback transcripts, notes) back into the canvas."` + "\n")
	b.WriteString("---\n\n")
	b.WriteString("# cairn — Product Design Thinking method\n\n")
	b.WriteString("Five stages on the EDT Loop. Each stage reads its reference file, writes ONE markdown artefact, prints a playback summary, and STOPS — the pause is the point. Never chain into the next stage automatically.\n\n")
	b.WriteString("| Stage | Loop | Reference | Artefact |\n")
	b.WriteString("| --- | --- | --- | --- |\n")
	for _, t := range tmpls {
		b.WriteString("| " + t.Name + " | " + t.Stage + " | references/" + t.ID + ".md | " + t.Description + " |\n")
	}
	b.WriteString("\nResolve artefact paths from the repo's `.cairn.yaml` if present (see any reference file's Config section); otherwise use the kit-native `design/` tree.\n")
	b.WriteString("\nAfter finishing a stage, tell the user where they are in the Loop and what the legitimate next moves are (the `next:` line of the reference file).\n")

	return writeBundleFile(filepath.Join(skillDir, "SKILL.md"), []byte(b.String()), force)
}

func writeBundleFile(path string, content []byte, force bool) GenerateResult {
	if _, err := os.Stat(path); err == nil && !force {
		return GenerateResult{
			Success:  false,
			FilePath: path,
			Message:  "file already exists (use --force to overwrite)",
			Error:    internal.ErrFileExists,
		}
	}

	if err := os.WriteFile(path, content, 0o644); err != nil {
		return GenerateResult{
			Success:  false,
			FilePath: path,
			Message:  "failed to write file: " + err.Error(),
			Error:    err,
		}
	}

	return GenerateResult{
		Success:  true,
		FilePath: path,
		Message:  "template generated successfully",
	}
}

func stripFrontmatter(content string) string {
	if !strings.HasPrefix(content, "---") {
		return content
	}
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return content
	}
	return strings.TrimLeft(parts[2], "\n")
}
