package model

import (
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
	imageText     string
	imageLength   int
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
		imageText:     image.Image2Ascii(),
		imageLength:   30,
	}
}
func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) View() string {

	return lipgloss.JoinHorizontal(lipgloss.Center, m.imageViewPort.View(), "  ", m.textViewport.View())

}
func (m *Model) sizeInputs() {

	m.textViewport.Width = m.width - m.imageLength
	m.textViewport.Height = m.height

	m.imageViewPort.Height = m.height
	m.imageViewPort.Width = m.imageLength
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
	)

	m.textViewport, vpCmd = m.textViewport.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// Set size of window
		m.height = msg.Height
		m.width = msg.Width

		// Get info to show
		textContent := GetInfoString(m)

		// Wrap content before setting it.
		m.textViewport.SetContent(textContent)
		m.imageViewPort.SetContent(m.imageText)

		m.textViewport.GotoBottom()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+z":
			return m, tea.Suspend
		}
	}
	m.sizeInputs()

	return m, tea.Batch(vpCmd)
}
