package model

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tekofx/fursona-tui/internal/config"
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
