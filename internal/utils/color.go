package utils

import (
	"strconv"
	"strings"
)

// GetContrastColor returns "#000000" (black) or "#FFFFFF" (white) for best contrast
func GetContrastColor(hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		// Expand shorthand hex (e.g., "abc" -> "aabbcc")
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}
	if len(hex) != 6 {
		return "#000000" // fallback to black if invalid
	}

	r, _ := strconv.ParseInt(hex[0:2], 16, 64)
	g, _ := strconv.ParseInt(hex[2:4], 16, 64)
	b, _ := strconv.ParseInt(hex[4:6], 16, 64)

	// Calculate luminance (per ITU-R BT.709)
	luminance := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)

	if luminance > 128 {
		return "#000000" // black for light backgrounds
	}
	return "#FFFFFF" // white for dark backgrounds
}
