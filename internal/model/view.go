package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	if !m.ready {
		return ""
	}
	if m.width < minWidth || m.height < minHeight {
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
