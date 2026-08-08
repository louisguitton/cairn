package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// FileName is the committed config file cairn looks for.
const FileName = ".cairn.yaml"

// RepoRole binds one of the two repo roles of the cairn method.
type RepoRole struct {
	// Path is relative to the repo holding .cairn.yaml, or absolute.
	// "." means this repo. Empty means the role is unbound.
	Path string `yaml:"path,omitempty"`
}

// Bindings maps artefact types onto the host repo's existing structure.
// Every field is optional; unbound types fall back to the kit-native tree.
// Paths are relative to the docs repo (or this repo when docs_repo is unbound).
type Bindings struct {
	Hills     string `yaml:"hills,omitempty"`     // e.g. product/hills.md (macro hill register, read-only)
	Personas  string `yaml:"personas,omitempty"`  // e.g. product/personas.md
	Specs     string `yaml:"specs,omitempty"`     // e.g. product/specs (per-feature canvas folders)
	Capture   string `yaml:"capture,omitempty"`   // e.g. stream (dated intake briefs)
	Decisions string `yaml:"decisions,omitempty"` // e.g. decisions (ADR promotion target)
	Glossary  string `yaml:"glossary,omitempty"`  // e.g. glossary.md
}

// Gating declares how a repo expects artefacts to be shared once a human has
// reviewed them. It is descriptive, not an authorisation: stage commands write
// artefacts to the working tree and never commit, push, or open a PR on their
// own.
type Gating struct {
	// Mode: "pr" (share via pull/merge request), "commit" (commit locally),
	// or "none" (leave the file in the working tree). Default "none".
	Mode string `yaml:"mode,omitempty"`
	// Frontmatter is an optional path to a frontmatter template that must
	// be prepended to every truth artefact (host-repo contract).
	Frontmatter string `yaml:"frontmatter,omitempty"`
}

// Config is the parsed .cairn.yaml.
type Config struct {
	Version   int      `yaml:"version"`
	DocsRepo  RepoRole `yaml:"docs_repo,omitempty"`
	ProtoRepo RepoRole `yaml:"proto_repo,omitempty"`
	Bindings  Bindings `yaml:"bindings,omitempty"`
	Gating    Gating   `yaml:"gating,omitempty"`

	// dir is the directory .cairn.yaml was loaded from (not serialized).
	dir string
}

// ArtefactType enumerates what Resolve can be asked for.
type ArtefactType string

const (
	Hills     ArtefactType = "hills"
	Personas  ArtefactType = "personas"
	Specs     ArtefactType = "specs"
	Capture   ArtefactType = "capture"
	Decisions ArtefactType = "decisions"
	Glossary  ArtefactType = "glossary"
)

// AllArtefactTypes in stable display order.
var AllArtefactTypes = []ArtefactType{Hills, Personas, Specs, Capture, Decisions, Glossary}

// kit-native fallback paths, relative to the repo running the command.
var fallbacks = map[ArtefactType]string{
	Hills:     "design/hill",
	Personas:  "design/personas.md",
	Specs:     "design/canvas",
	Capture:   "design/intake",
	Decisions: "design/decisions",
	Glossary:  "design/glossary.md",
}

// Resolution is one row of the path resolution table.
type Resolution struct {
	Type     ArtefactType
	Path     string // absolute
	Bound    bool   // true when an explicit binding was used
	Exists   bool   // whether the path currently exists on disk
	ReadOnly bool   // hills macro register is read-only from micro flows
}

// Load searches dir and its parents for .cairn.yaml. A missing file is not
// an error: it returns a zero-value Config rooted at dir (all fallbacks).
func Load(dir string) (Config, error) {
	abs, err := filepath.Abs(dir)
	if err != nil {
		return Config{}, err
	}

	for d := abs; ; d = filepath.Dir(d) {
		candidate := filepath.Join(d, FileName)
		if _, err := os.Stat(candidate); err == nil {
			data, err := os.ReadFile(candidate)
			if err != nil {
				return Config{}, fmt.Errorf("read %s: %w", candidate, err)
			}
			var cfg Config
			if err := yaml.Unmarshal(data, &cfg); err != nil {
				return Config{}, fmt.Errorf("parse %s: %w", candidate, err)
			}
			cfg.dir = d
			return cfg, nil
		}
		if filepath.Dir(d) == d {
			break
		}
	}

	return Config{Version: 1, dir: abs}, nil
}

// Dir returns the directory the config is rooted at.
func (c Config) Dir() string { return c.dir }

// docsRoot returns the absolute directory bindings resolve against.
func (c Config) docsRoot() string {
	p := c.DocsRepo.Path
	if p == "" || p == "." {
		return c.dir
	}
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(c.dir, p)
}

func (c Config) binding(t ArtefactType) string {
	switch t {
	case Hills:
		return c.Bindings.Hills
	case Personas:
		return c.Bindings.Personas
	case Specs:
		return c.Bindings.Specs
	case Capture:
		return c.Bindings.Capture
	case Decisions:
		return c.Bindings.Decisions
	case Glossary:
		return c.Bindings.Glossary
	default:
		return ""
	}
}

// Resolve returns where an artefact type lives: the explicit binding
// (relative to the docs repo) when set, else the kit-native fallback
// (relative to the config root).
func (c Config) Resolve(t ArtefactType) Resolution {
	if bound := c.binding(t); bound != "" {
		path := bound
		if !filepath.IsAbs(path) {
			path = filepath.Join(c.docsRoot(), bound)
		}
		return c.finish(Resolution{Type: t, Path: path, Bound: true})
	}
	return c.finish(Resolution{Type: t, Path: filepath.Join(c.dir, fallbacks[t]), Bound: false})
}

func (c Config) finish(r Resolution) Resolution {
	if _, err := os.Stat(r.Path); err == nil {
		r.Exists = true
	}
	r.ReadOnly = r.Type == Hills && r.Bound
	return r
}

// ResolveAll returns the full resolution table in stable order.
func (c Config) ResolveAll() []Resolution {
	out := make([]Resolution, 0, len(AllArtefactTypes))
	for _, t := range AllArtefactTypes {
		out = append(out, c.Resolve(t))
	}
	return out
}

// Save writes the config as .cairn.yaml into dir.
func Save(cfg Config, dir string) (string, error) {
	if cfg.Version == 0 {
		cfg.Version = 1
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return "", err
	}
	header := []byte("# cairn — prompt-driven product design thinking kit\n# Repo roles and artefact path bindings. Docs: https://github.com/louisguitton/cairn\n")
	path := filepath.Join(dir, FileName)
	if err := os.WriteFile(path, append(header, data...), 0o644); err != nil {
		return "", err
	}
	return path, nil
}
