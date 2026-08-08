package detector

// AIToolType represents the type of AI assistant tool cairn can target.
type AIToolType string

const (
	ClaudeCode AIToolType = "claude-code"
	Cursor     AIToolType = "cursor"
	// Claude is the claude.ai / Claude Desktop target. It has no repo
	// signature and is never auto-detected: it is only reachable via
	// --tool claude or the interactive picker. Generation emits a skill
	// bundle for manual upload instead of a commands directory.
	Claude  AIToolType = "claude"
	Unknown AIToolType = "unknown"
)

// String returns the human-readable name of the tool type.
func (t AIToolType) String() string {
	switch t {
	case ClaudeCode:
		return "Claude Code"
	case Cursor:
		return "Cursor"
	case Claude:
		return "Claude"
	default:
		return "Unknown"
	}
}

// GetConfigDir returns the directory (relative to the working dir) that
// generated templates are written to for each tool type.
func (t AIToolType) GetConfigDir() string {
	switch t {
	case ClaudeCode:
		return ".claude/commands"
	case Cursor:
		return ".cursor/commands"
	case Claude:
		return "cairn-skill"
	default:
		return ""
	}
}

// GetSignatureFiles returns the list of signature files/directories used for
// auto-detection. Claude returns nil on purpose — see the type comment.
func (t AIToolType) GetSignatureFiles() []string {
	switch t {
	case ClaudeCode:
		return []string{".claude", "CLAUDE.md"}
	case Cursor:
		return []string{".cursor", ".cursorrules"}
	default:
		return nil
	}
}

// DetectResult holds the result of environment detection.
type DetectResult struct {
	ToolType   AIToolType
	ConfigPath string
	IsValid    bool
	Message    string
}
