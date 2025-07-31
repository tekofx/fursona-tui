package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tekofx/fursona-tui/internal/model"
)

func main() {

	// Initialize our program
	p := tea.NewProgram(model.InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}
