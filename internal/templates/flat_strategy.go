package templates

import (
	"path/filepath"

	"github.com/louisguitton/cairn/internal/detector"
)

type FlatMarkdownStrategy struct {
	tool    detector.AIToolType
	manager *EmbeddedTemplateManager
}

func (s *FlatMarkdownStrategy) GenerateAll(workingDir string, force bool) []GenerateResult {
	tmpls, err := s.manager.ListAvailable(s.tool)
	if err != nil {
		return []GenerateResult{{
			Success: false,
			Message: "failed to list templates: " + err.Error(),
			Error:   err,
		}}
	}

	results := make([]GenerateResult, 0, len(tmpls))
	for _, t := range tmpls {
		results = append(results, s.GenerateOne(workingDir, t, force)...)
	}
	return results
}

func (s *FlatMarkdownStrategy) GenerateOne(workingDir string, tmpl TemplateMeta, force bool) []GenerateResult {
	targetDir := workingDir
	if configDir := s.tool.GetConfigDir(); configDir != "" {
		targetDir = filepath.Join(workingDir, configDir)
	}

	targetPath := filepath.Join(targetDir, tmpl.ID+".md")
	req := GenerateRequest{
		TemplateName: tmpl.Name,
		TargetPath:   targetPath,
		Force:        force,
	}
	return []GenerateResult{s.manager.Generate(req)}
}
