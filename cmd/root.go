package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/louisguitton/cairn/internal/detector"
	"github.com/louisguitton/cairn/internal/templates"
	"github.com/louisguitton/cairn/internal/ui"
)

var (
	det             detector.DetectorService
	uiRenderer      ui.UIRenderer
	templateManager templates.TemplateManager
	detectedResult  detector.DetectResult
	toolFlag        string
)

var rootCmd = &cobra.Command{
	Use:   "cairn",
	Short: "Prompt-driven Product Design Thinking kit on the IBM EDT Loop",
	Long: `cairn installs a five-stage, gated Product Design Thinking pipeline
(intake → hill → canvas → handoff → sync) as agent command templates.

Supports Claude Code, Cursor, and Claude (skill bundle).
Auto-detects your environment and manages the stage templates.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		det = detector.NewDefaultDetector()
		uiRenderer = ui.NewCharmUIRenderer()
		templateManager = templates.NewEmbeddedTemplateManager()

		workingDir, _ := os.Getwd()

		if toolFlag != "" {
			toolType := ParseToolFlag(toolFlag)
			detectedResult = detector.DetectResult{
				ToolType:   toolType,
				ConfigPath: det.GetConfigDirPath(toolType, workingDir),
				IsValid:    toolType != detector.Unknown,
				Message:    "tool manually specified: " + toolType.String(),
			}
		} else {
			detectedResult = det.Detect(workingDir)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&toolFlag, "tool", "t", "", "Manually specify tool type (claude-code, cursor, claude)")
}

func SetVersion(v string) {
	v = strings.TrimSpace(v)
	if v == "" {
		v = "dev"
	}
	rootCmd.Version = v
	rootCmd.SetVersionTemplate("cairn {{.Version}}\n")

	// Cobra registers the `--version` flag lazily when Version is non-empty
	// and the command is executed. To bind a `-v` shorthand we touch the
	// flag explicitly: declare it ourselves if absent, otherwise patch the
	// existing definition.
	if f := rootCmd.Flags().Lookup("version"); f != nil {
		f.Shorthand = "v"
		f.Usage = "Print cairn version and exit"
	} else {
		rootCmd.Flags().BoolP("version", "v", false, "Print cairn version and exit")
	}
}

// RootCommand exposes the root *cobra.Command for testing. Production code
// should use Execute() instead.
func RootCommand() *cobra.Command {
	return rootCmd
}

// Execute runs the root command.
func Execute() {
	maybePrintPathHint()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// ParseToolFlag converts a tool flag string to AIToolType.
func ParseToolFlag(flag string) detector.AIToolType {
	switch strings.ToLower(flag) {
	case "claude-code", "claudecode", "cc":
		return detector.ClaudeCode
	case "cursor":
		return detector.Cursor
	case "claude":
		return detector.Claude
	default:
		return detector.Unknown
	}
}
