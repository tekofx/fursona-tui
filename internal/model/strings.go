package model

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/tekofx/fursona-tui/internal/style"
	"github.com/tekofx/fursona-tui/internal/utils"
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

// Returns the data formatted as strings
func GetStrings(m Model) string {
	output := ""

	output += getInfoString(m)
	if m.Config.Quote != "" {
		output += fmt.Sprintf("\n%s\n", getQuote(m))
	}
	if len(m.Config.Palette) != 0 {
		output += fmt.Sprintf("\n%s \n", getColorPalette(m))
	}

	return output
}

func getQuote(m Model) string {
	return style.Quote.Render(m.Config.Quote)
}

func getInfoString(m Model) string {
	infoString := ""
	infoString += style.Heading1.Render(m.Config.Name)

	infoString += "\n--------\n" // Add OtherData key-value pairs
	if len(m.Config.Data) > 0 {
		for k, v := range m.Config.Data {

			infoString += fmt.Sprintf("%s: %s\n", style.Key.Render(k), style.Default.Render(v))
		}
	}

	return infoString
}
