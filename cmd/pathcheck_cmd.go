package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/louisguitton/cairn/internal/config"
)

var pathcheckCmd = &cobra.Command{
	Use:   "pathcheck",
	Short: "Print the artefact path resolution table from .cairn.yaml",
	Long: `Shows, for every artefact type, where cairn commands will read/write:
the explicit binding from .cairn.yaml when set, else the kit-native fallback.
Dangling paths are flagged as warnings, never errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		workingDir, _ := os.Getwd()
		cfg, err := config.Load(workingDir)
		if err != nil {
			uiRenderer.RenderError("Failed to load " + config.FileName + ": " + err.Error())
			return
		}

		var rows [][]string
		var dangling int
		for _, r := range cfg.ResolveAll() {
			source := "fallback"
			if r.Bound {
				source = "bound"
			}
			exists := "missing"
			if r.Exists {
				exists = "ok"
			} else if r.Bound {
				dangling++
			}
			mode := "rw"
			if r.ReadOnly {
				mode = "read-only"
			}
			rows = append(rows, []string{string(r.Type), r.Path, source, exists, mode})
		}

		uiRenderer.RenderTable([]string{"Artefact", "Path", "Source", "Status", "Mode"}, rows)

		if dangling > 0 {
			uiRenderer.RenderWarning("Some bound paths do not exist yet — check .cairn.yaml or create them")
		}
		if cfg.Gating.Mode != "" {
			uiRenderer.RenderSuccess("Gating mode: " + cfg.Gating.Mode)
		}
	},
}

func init() {
	rootCmd.AddCommand(pathcheckCmd)
}
