package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/style"
	"github.com/tekofx/fursona-tui/utils"
)

func getColorPalette(m Model) string {

	output := style.Default.Render("Color Palette")

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

	infoString += style.Default.Render(m.Name)
	if m.Surname != "" {
		infoString += fmt.Sprintf(" %s\n", style.Default.Render(m.Surname))
	}
	infoString += "--------\n"
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Species"), style.Default.Render(m.Species))
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Gender"), style.Default.Render(m.Gender))
	infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Pronouns"), style.Default.Render(m.Pronouns))
	infoString += fmt.Sprintf("\n%s \n", getColorPalette(m))
	return infoString
}
