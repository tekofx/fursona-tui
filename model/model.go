package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/config"
	"github.com/tekofx/fursona-tui/image"
)

type Model struct {
	width         int
	height        int
	Name          string
	Surname       string
	Species       string
	Gender        string
	Pronouns      string
	Palette       []string
	textViewport  viewport.Model
	imageViewPort viewport.Model
}

func InitialModel() Model {
	textViewport := viewport.New(150, 50)
	imageViewport := viewport.New(50, 50)
	config := config.ReadConfig()
	return Model{
		Name:          config.Name,
		Surname:       config.Surname,
		Species:       config.Species,
		Gender:        config.Gender,
		Pronouns:      config.Pronouns,
		Palette:       config.Palette,
		textViewport:  textViewport,
		imageViewPort: imageViewport,
	}
}
func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) View() string {

	return lipgloss.JoinHorizontal(lipgloss.Top, m.imageViewPort.View(), m.textViewport.View())

}
func (m *Model) sizeInputs() {

	m.textViewport.Width = m.width / 2
	m.textViewport.Height = m.height

	m.imageViewPort.Width = m.width / 2
	m.imageViewPort.Height = m.height
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
	)

	m.textViewport, vpCmd = m.textViewport.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		fmt.Print(msg.Width)
		// Set size of window
		m.height = msg.Height
		m.width = msg.Width
		m.sizeInputs()

		// Get info to show
		textContent := GetInfoString(m)
		imageContent := image.Image2Ascii()

		// Wrap content before setting it.
		m.textViewport.SetContent(textContent)
		m.imageViewPort.SetContent(imageContent)

		m.textViewport.GotoBottom()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}
	}

	return m, tea.Batch(vpCmd)
}
