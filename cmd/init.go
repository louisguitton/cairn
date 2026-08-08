package cmd

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"

	"github.com/louisguitton/cairn/internal/config"
	"github.com/louisguitton/cairn/internal/detector"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize cairn in this repo: declare repo roles, bind artefact paths, create the command directory",
	Long: `Interactive setup for the cairn method in an existing project.

Declares this repo's role (prototype / docs / both), optionally points at a
docs repo, confirms artefact path bindings (hills, personas, specs, capture),
writes .cairn.yaml, and creates the detected AI tool's command directory.

Nothing is bound silently: every auto-suggested path must be confirmed.`,
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, _ := os.Getwd()

		if !detectedResult.IsValid {
			tool := selectToolInteractively()
			if tool == detector.Unknown {
				uiRenderer.RenderError("No tool selected")
				return
			}
			detectedResult = detector.DetectResult{
				ToolType:   tool,
				ConfigPath: det.GetConfigDirPath(tool, workingDir),
				IsValid:    true,
				Message:    "tool manually selected: " + tool.String(),
			}
		}

		cfg, ok := runInitInterview(workingDir)
		if !ok {
			uiRenderer.RenderWarning("Initialization cancelled")
			return
		}

		path, err := config.Save(cfg, workingDir)
		if err != nil {
			uiRenderer.RenderError("Failed to write " + config.FileName + ": " + err.Error())
			return
		}
		uiRenderer.RenderSuccess("Wrote " + path)

		configPath := detectedResult.ConfigPath
		if configPath != "" {
			if err := os.MkdirAll(configPath, 0o755); err != nil {
				uiRenderer.RenderError("Failed to create directory: " + err.Error())
				return
			}
			uiRenderer.RenderSuccess("Initialized " + detectedResult.ToolType.String() + " command directory at: " + configPath)
		}

		uiRenderer.RenderSuccess("Next: run `cairn generate --all` to install the stage commands, then `cairn pathcheck` to review artefact paths")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// runInitInterview collects repo roles and bindings. Returns ok=false on cancel.
func runInitInterview(workingDir string) (config.Config, bool) {
	cfg := config.Config{Version: 1}

	var role string
	roleForm := huh.NewForm(huh.NewGroup(
		huh.NewSelect[string]().
			Title("What is the role of this repo in the cairn method?").
			Options(
				huh.NewOption("Prototype repo (Make happens here; docs live elsewhere)", "proto"),
				huh.NewOption("Docs repo / second brain (truth artefacts live here)", "docs"),
				huh.NewOption("Both (solo repo: docs and prototype together)", "both"),
			).
			Value(&role),
	))
	if err := roleForm.Run(); err != nil {
		return cfg, false
	}

	switch role {
	case "proto":
		cfg.ProtoRepo.Path = "."
		var docsPath string
		docsForm := huh.NewForm(huh.NewGroup(
			huh.NewInput().
				Title("Path to the docs repo (empty = none yet; artefacts fall back to design/ here)").
				Placeholder("../team-docs").
				Value(&docsPath),
		))
		if err := docsForm.Run(); err != nil {
			return cfg, false
		}
		if docsPath != "" {
			cfg.DocsRepo.Path = docsPath
		}
	case "docs":
		cfg.DocsRepo.Path = "."
	case "both":
		cfg.DocsRepo.Path = "."
		cfg.ProtoRepo.Path = "."
	}

	docsRoot := workingDir
	if cfg.DocsRepo.Path != "" && cfg.DocsRepo.Path != "." {
		if filepath.IsAbs(cfg.DocsRepo.Path) {
			docsRoot = cfg.DocsRepo.Path
		} else {
			docsRoot = filepath.Join(workingDir, cfg.DocsRepo.Path)
		}
	}

	if cfg.DocsRepo.Path != "" {
		cfg.Bindings = interviewBindings(docsRoot)

		var gatingMode string
		gatingForm := huh.NewForm(huh.NewGroup(
			huh.NewSelect[string]().
				Title("How are truth artefact writes gated in the docs repo?").
				Options(
					huh.NewOption("PR — branch + pull request, merge is the playback sign-off", "pr"),
					huh.NewOption("Commit — direct commits allowed", "commit"),
					huh.NewOption("None — plain file writes", "none"),
				).
				Value(&gatingMode),
		))
		if err := gatingForm.Run(); err != nil {
			return cfg, false
		}
		cfg.Gating.Mode = gatingMode
	}

	return cfg, true
}

// interviewBindings probes the docs repo for likely artefact homes and asks
// the human to confirm each suggestion. Never binds silently (Safeguard 8).
func interviewBindings(docsRoot string) config.Bindings {
	b := config.Bindings{}
	b.Hills = confirmBinding(docsRoot, "hills (macro hill register, linked read-only)",
		[]string{"product/hills.md", "docs/hills.md", "hills.md"})
	b.Personas = confirmBinding(docsRoot, "personas",
		[]string{"product/personas.md", "docs/personas.md", "personas.md"})
	b.Specs = confirmBinding(docsRoot, "specs (per-feature canvas folders)",
		[]string{"product/specs", "docs/specs", "specs"})
	b.Capture = confirmBinding(docsRoot, "capture (dated intake briefs, append-only)",
		[]string{"stream", "capture", "inbox"})
	b.Decisions = confirmBinding(docsRoot, "decisions (ADR promotion target)",
		[]string{"decisions", "docs/decisions", "docs/architecture"})
	b.Glossary = confirmBinding(docsRoot, "glossary",
		[]string{"glossary.md", "docs/glossary.md"})
	return b
}

// confirmBinding suggests the first candidate that exists under docsRoot and
// lets the user confirm, edit, or clear it. Returns "" for unbound.
func confirmBinding(docsRoot, label string, candidates []string) string {
	suggestion := ""
	for _, c := range candidates {
		if _, err := os.Stat(filepath.Join(docsRoot, c)); err == nil {
			suggestion = c
			break
		}
	}

	value := suggestion
	form := huh.NewForm(huh.NewGroup(
		huh.NewInput().
			Title("Binding for " + label + " (empty = unbound, kit-native fallback)").
			Value(&value),
	))
	if err := form.Run(); err != nil {
		return ""
	}
	return value
}

func selectToolInteractively() detector.AIToolType {
	options := []huh.Option[string]{
		huh.NewOption("Claude Code", "claude-code"),
		huh.NewOption("Cursor", "cursor"),
		huh.NewOption("Claude (claude.ai / Desktop — emits a skill bundle)", "claude"),
	}

	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select your AI tool").
				Options(options...).
				Value(&selected),
		),
	)

	if err := form.Run(); err != nil {
		return detector.Unknown
	}

	return ParseToolFlag(selected)
}
