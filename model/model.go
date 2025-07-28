package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/image"
	"github.com/tekofx/fursona-tui/style"
)

const gap = "\n\n"

type Model struct {
	width         int
	height        int
	name          string
	surname       string
	species       string
	textViewport  viewport.Model
	imageViewPort viewport.Model
}

func InitialModel() Model {
	textViewport := viewport.New(30, 5)
	imageViewport := viewport.New(30, 30)

	return Model{
		name:          "Teko",
		surname:       "Fresnes Xaiden",
		species:       "Arctic Fox",
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
		m.height = msg.Height
		m.width = msg.Width
		textContent := fmt.Sprintf("%s\n%s\n%s",
			style.H1.Render(m.name),
			style.Dimmed.Render(m.surname),
			style.Dimmed.Render(m.species),
		)

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
	m.sizeInputs()

	return m, tea.Batch(vpCmd)
}
