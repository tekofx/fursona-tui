package model

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/qeesung/image2ascii/convert"
	"github.com/tekofx/fursona-tui/style"
)

const gap = "\n\n"

func image2Ascii() string {
	converter := convert.NewImageConverter()
	ascii := converter.ImageFile2ASCIIString("logo2.png", &convert.Options{
		Colored:    true,
		Ratio:      1.0,
		FixedWidth: 50,
	})
	return ascii
}

type Model struct {
	name     string
	surname  string
	species  string
	viewport viewport.Model
}

func InitialModel() Model {
	vp := viewport.New(30, 5)
	vp.SetContent(`Welcome to the chat room!
			Type a message and press Enter to send.`)

	return Model{
		name:     "Teko",
		surname:  "Fresnes Xaiden",
		species:  "Arctic Fox",
		viewport: vp,
	}
}
func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) View() string {
	return fmt.Sprintf(
		"%s",
		m.viewport.View(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
	)

	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width

		content := fmt.Sprintf("%s\n%s\n%s",
			style.H1.Render(m.name),
			style.Dimmed.Render(m.surname),
			style.Dimmed.Render(m.species),
		)

		// Wrap content before setting it.
		m.viewport.SetContent(content)

		m.viewport.GotoBottom()

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
