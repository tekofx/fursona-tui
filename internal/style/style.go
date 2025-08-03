package style

import "github.com/charmbracelet/lipgloss"

var (
	Default  = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	Key      = lipgloss.NewStyle().Foreground(lipgloss.Color("#EB58C7"))
	Error    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4D4D"))
	Heading1 = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
	Heading2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Underline(true)
	Quote    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Italic(true)
)
