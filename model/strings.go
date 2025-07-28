package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/style"
	"github.com/tekofx/fursona-tui/utils"
)

func getColorPalette(m Model) string {

	output := style.H1.Render("Color Palette")

	for i, color := range m.Palette {
		if i%5 == 0 {
			output += "\n"
		}
		text := utils.GetContrastColor(color)
		output += lipgloss.NewStyle().Background(lipgloss.Color(color)).Foreground(lipgloss.Color(text)).PaddingLeft(1).PaddingRight(1).Render(color)
	}

	return output

}

func GetInfoString(m Model) string {
	infoString := ""

	infoString += style.H1.Render(m.Name)
	if m.Surname != "" {
		infoString += fmt.Sprintf(" %s\n", style.H1.Render(m.Surname))
	}
	infoString += "--------\n"
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Species"), style.Dimmed.Render(m.Species))
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Gender"), style.Dimmed.Render(m.Gender))
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Pronouns"), style.Dimmed.Render(m.Pronouns))
	infoString += fmt.Sprintf("\n%s \n", getColorPalette(m))
	return infoString
}
