package model

import (
	"fmt"
	"image/png"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/nfnt/resize"
	"github.com/qeesung/image2ascii/convert"
	"github.com/tekofx/fursona-tui/style"
)

const gap = "\n\n"

func image2Ascii() string {

	file, _ := os.Open("logo2.png")
	img, _ := png.Decode(file)

	// Get original dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	// Aspect ratio correction factor (experiment with 0.5–0.6)
	aspectRatio := 0.5
	newHeight := uint(float64(height) * aspectRatio)

	// Resize image
	resized := resize.Resize(uint(width), newHeight, img, resize.Lanczos3)

	converter := convert.NewImageConverter()
	ascii := converter.Image2ASCIIString(resized, &convert.Options{
		Colored:     true,
		FixedHeight: 20,
	})
	return ascii
}

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

		imageContent := image2Ascii()

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
