package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tekofx/fursona-tui/internal/image"
)

func (m *Model) sizeInputs() {

	m.textViewport.Width = (m.width / 2)
	m.textViewport.Height = m.height - m.verticalPadding

	m.imageViewPort.Width = (m.width / 2)
	m.imageViewPort.Height = m.height - m.verticalPadding
}
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		// Set size of window
		m.height = msg.Height
		m.width = msg.Width

		if m.width > minWidth && m.height > minHeight {
			// Get info to show
			textContent := GetStrings(m)

			// Wrap content before setting it.
			m.textViewport.SetContent(textContent)
			m.imageViewPort.SetContent(image.Image2Ascii((msg.Width / 2) - 2))

			m.textViewport.GotoBottom()
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			if m.help.ShowAll {
				m.verticalPadding = 1
			} else {
				m.verticalPadding = 2
			}
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}
	m.sizeInputs()

	return m, nil
}
