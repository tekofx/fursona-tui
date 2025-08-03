package model

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tekofx/fursona-tui/internal/image"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:

		// Set size of window
		m.height = msg.Height
		m.width = msg.Width

		m.textViewport.GotoBottom()
		if !m.ready {
			m.textViewport = viewport.New(m.width, m.height-m.verticalPadding)
			m.imageViewPort = viewport.New(m.width/2, m.height-m.verticalPadding)
			m.textViewport.SetContent(GetStrings(m))
			m.imageViewPort.SetContent(image.Image2Ascii((msg.Width / 2) - 2))
			m.ready = true
		} else {
			m.textViewport.Width = (m.width)
			m.textViewport.Height = m.height - m.verticalPadding

			m.imageViewPort.Width = (m.width / 2)
			m.imageViewPort.Height = m.height - m.verticalPadding
			m.imageViewPort.SetContent(image.Image2Ascii((msg.Width / 2) - 2))

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

	return m, nil
}
