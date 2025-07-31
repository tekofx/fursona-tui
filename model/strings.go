package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/style"
	"github.com/tekofx/fursona-tui/utils"
)

func getColorPalette(m Model) string {

	output := style.Heading2.Render("Color Palette")

	for i, color := range m.Config.Palette {
		if i%5 == 0 {
			output += "\n"
		}
		text := utils.GetContrastColor(color)
		output += lipgloss.
			NewStyle().
			Background(lipgloss.Color(color)).
			Foreground(lipgloss.Color(text)).
			PaddingLeft(1).
			PaddingRight(1).
			Render(color)
	}

	return output

}

func GetInfoString(m Model) string {
	infoString := ""
	infoString += style.Heading1.Render(m.Config.Name)
	if m.Config.Surname != "" {
		infoString += fmt.Sprintf(" %s\n", style.Heading1.Render(m.Config.Surname))
	}
	infoString += "--------\n"

	if m.Config.Species != "" {
		infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Species"), style.Default.Render(m.Config.Species))
	}
	if m.Config.Gender != "" {
		infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Gender"), style.Default.Render(m.Config.Gender))
	}
	if m.Config.Pronouns != "" {
		infoString += fmt.Sprintf("%s: %s\n", style.Key.Render("Pronouns"), style.Default.Render(m.Config.Pronouns))
	}
	if len(m.Config.Palette) != 0 {
		infoString += fmt.Sprintf("\n%s \n", getColorPalette(m))
	}

	// Add OtherData key-value pairs
	if len(m.Config.OtherData) > 0 {
		infoString += "\n" + style.Heading2.Render("Other Info") + "\n"
		for k, v := range m.Config.OtherData {
			infoString += fmt.Sprintf("%s: %s\n", style.Key.Render(k), style.Default.Render(v))
		}
	}

	return infoString
}
