package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	name    string
	surname string
	species string
}

var (
	defaultStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
)

func InitialModel() Model {
	return Model{
		name:    "Name",
		surname: "Surname",
		species: "Species",
	}
}
func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) View() string {
	msg := defaultStyle.Render("Hello")
	return fmt.Sprintf(msg)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}
	}

	return m, nil
}
