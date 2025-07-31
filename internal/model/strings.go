package model

import (
	"fmt"
	"strings"

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

			infoString += fmt.Sprintf("%s: %s\n", style.Key.Render(HumanizeKey(k)), style.Default.Render(v))
		}
	}

	return infoString
}
func HumanizeKey(key string) string {
	parts := strings.Split(key, "_")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(string(part[0])) + strings.ToLower(part[1:])
		}
	}
	return strings.Join(parts, " ")
}
