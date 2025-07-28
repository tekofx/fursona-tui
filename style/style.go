package style

import "github.com/charmbracelet/lipgloss"

var (
	Default = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	H1      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
	Dimmed  = lipgloss.NewStyle().Foreground(lipgloss.Color("#A6A6A6"))
)
