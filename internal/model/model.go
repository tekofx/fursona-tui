package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/internal/config"
	"github.com/tekofx/fursona-tui/internal/image"
)

const minWidth = 100
const minHeight = 25

type Model struct {
	keys          keyMap
	help          help.Model
	width         int
	height        int
	Config        config.Config
	textViewport  viewport.Model
	imageViewPort viewport.Model
}

type keyMap struct {
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},
	}
}

var keys = keyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

func InitialModel() Model {
	textViewport := viewport.New(150, 50)
	imageViewport := viewport.New(50, 50)
	cfg := config.ReadConfig()
	return Model{
		keys:          keys,
		help:          help.New(),
		Config:        *cfg,
		textViewport:  textViewport,
		imageViewPort: imageViewport,
	}
}
func (m Model) Init() tea.Cmd {

	return nil
}

func (m Model) View() string {
	if (m.width != 0 && m.width < minWidth) || (m.height != 0 && m.height < minHeight) {
		// Center the message using Lipgloss
		msg := fmt.Sprintf("Window too small!\nPlease resize your terminal.\n%dx%d", m.width, m.height)
		style := lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Align(lipgloss.Center)
		return style.Render(msg)
	}
	horizontal := lipgloss.JoinHorizontal(lipgloss.Center, " ", m.imageViewPort.View(), "  ", m.textViewport.View(), " ")
	return lipgloss.JoinVertical(lipgloss.Center, "\n", horizontal, m.help.View(m.keys))

}
func (m *Model) sizeInputs() {

	m.textViewport.Width = (m.width / 2) - 2
	m.textViewport.Height = m.height - 2

	m.imageViewPort.Width = (m.width / 2) - 2
	m.imageViewPort.Height = m.height - 2
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

		if m.width > minWidth && m.height > minHeight {
			// Get info to show
			textContent := GetInfoString(m)

			// Wrap content before setting it.
			m.textViewport.SetContent(textContent)
			m.imageViewPort.SetContent(image.Image2Ascii((msg.Width / 2) - 2))

			m.textViewport.GotoBottom()
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	m.sizeInputs()

	return m, tea.Batch(vpCmd)
}
