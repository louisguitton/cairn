package templates

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/louisguitton/cairn/internal"
	"github.com/louisguitton/cairn/internal/detector"
)

// TemplateManager defines the interface for template operations.
//
// Per-tool generation behavior (e.g., Claude's skill-bundle output) is
// intentionally NOT exposed on this interface.
// Use templates.StrategyFor(tool, mgr).GenerateAll(workingDir, force) instead.
type TemplateManager interface {
	ListCore() ([]TemplateMeta, error)
	ListBeta() ([]TemplateMeta, error)
	ListAvailable(tool detector.AIToolType) ([]TemplateMeta, error)
	ListAll() ([]TemplateMeta, error)
	GetByName(name string) (TemplateMeta, error)
	Generate(req GenerateRequest) GenerateResult
}

// EmbeddedTemplateManager implements TemplateManager using embedded templates.
type EmbeddedTemplateManager struct{}

// NewEmbeddedTemplateManager creates a new EmbeddedTemplateManager instance.
func NewEmbeddedTemplateManager() *EmbeddedTemplateManager {
	return &EmbeddedTemplateManager{}
}

// loadTemplatesFromDir loads and parses templates from a specific embedded directory path.
func (m *EmbeddedTemplateManager) loadTemplatesFromDir(dir string) ([]TemplateMeta, error) {
	entries, err := fs.ReadDir(embeddedTemplates, dir)
	if err != nil {
		return []TemplateMeta{}, nil
	}

	var templates []TemplateMeta
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}

		content, err := fs.ReadFile(embeddedTemplates, dir+"/"+entry.Name())
		if err != nil {
			continue
		}

		meta := ParseFrontmatter(string(content))
		if meta.ID == "" {
			meta.ID = strings.TrimSuffix(entry.Name(), ".md")
		}
		templates = append(templates, meta)
	}

	sort.Slice(templates, func(i, j int) bool {
		return templates[i].Name < templates[j].Name
	})

	return templates, nil
}

// ListCore returns all core templates that should be installed by default.
func (m *EmbeddedTemplateManager) ListCore() ([]TemplateMeta, error) {
	return m.loadTemplatesFromDir("data/core")
}

// ListBeta returns all beta templates available for manual selection.
func (m *EmbeddedTemplateManager) ListBeta() ([]TemplateMeta, error) {
	return m.loadTemplatesFromDir("data/beta")
}

// ListAvailable returns all templates installed by default for a tool.
// Every tool gets the same core set; beta templates are opt-in via list/generate.
func (m *EmbeddedTemplateManager) ListAvailable(tool detector.AIToolType) ([]TemplateMeta, error) {
	return m.ListCore()
}

// ListAll returns ALL templates across all categories (for admin/debug use).
func (m *EmbeddedTemplateManager) ListAll() ([]TemplateMeta, error) {
	coreTemplates, err := m.ListCore()
	if err != nil {
		return nil, err
	}

	templates := coreTemplates

	betaTemplates, err := m.ListBeta()
	if err == nil {
		templates = append(templates, betaTemplates...)
	}

	seen := make(map[string]bool)
	var unique []TemplateMeta
	for _, t := range templates {
		if !seen[t.ID] {
			seen[t.ID] = true
			unique = append(unique, t)
		}
	}

	sort.Slice(unique, func(i, j int) bool {
		return unique[i].Name < unique[j].Name
	})

	return unique, nil
}

// GetByName returns a template by its name (case-insensitive).
func (m *EmbeddedTemplateManager) GetByName(name string) (TemplateMeta, error) {
	templates, err := m.ListAll()
	if err != nil {
		return TemplateMeta{}, err
	}

	nameLower := strings.ToLower(name)
	for _, t := range templates {
		if strings.ToLower(t.Name) == nameLower || strings.ToLower(t.ID) == nameLower {
			return t, nil
		}
	}

	return TemplateMeta{}, internal.ErrTemplateNotFound
}

// Generate creates a template file at the specified target path.
func (m *EmbeddedTemplateManager) Generate(req GenerateRequest) GenerateResult {
	template, err := m.GetByName(req.TemplateName)
	if err != nil {
		return GenerateResult{
			Success: false,
			Message: "template not found: " + req.TemplateName,
			Error:   err,
		}
	}

	targetPath := req.TargetPath
	if targetPath == "" {
		return GenerateResult{
			Success: false,
			Message: "target path is required",
			Error:   fmt.Errorf("target path is required"),
		}
	}

	return m.generateWithContent(targetPath, template.Content, req.Force)
}

func (m *EmbeddedTemplateManager) generateWithContent(targetPath, content string, force bool) GenerateResult {
	if _, err := os.Stat(targetPath); err == nil && !force {
		return GenerateResult{
			Success:  false,
			FilePath: targetPath,
			Message:  "file already exists (use --force to overwrite)",
			Error:    internal.ErrFileExists,
		}
	}

	targetDir := filepath.Dir(targetPath)
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return GenerateResult{
			Success: false,
			Message: "failed to create directory: " + targetDir,
			Error:   fmt.Errorf("failed to create directory: %w", err),
		}
	}

	if err := os.WriteFile(targetPath, []byte(content), 0644); err != nil {
		return GenerateResult{
			Success: false,
			Message: "failed to write file: " + targetPath,
			Error:   fmt.Errorf("failed to write file: %w", err),
		}
	}

	return GenerateResult{
		Success:  true,
		FilePath: targetPath,
		Message:  "template generated successfully",
	}
}
