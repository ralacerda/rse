package app

import (
	"fmt"
	"strings"
)

func header(_ Model) string {
	return "What should we buy?"
}

func footer(_ Model) string {
	return "Press 'q' to quit"
}

func presets(m Model) string {
	// Check if any value from the presets match the current value
	p := m.findMatchingPreset()

	b := strings.Builder{}

	b.WriteString("Presets: \n")
	for i, preset := range m.presets {

		if p == preset.name {
			b.WriteString("✓")
		}

		b.WriteString(fmt.Sprintf("[%d] %s | ", i+1, preset.name))

	}

	return b.String()
}

func choices(m Model) string {
	b := strings.Builder{}

	for i, choice := range m.variables {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		value := choice.values[choice.selected]

		b.WriteString(fmt.Sprintf("%s [%s] %s // %s \n", cursor, value, choice.name, choice.description))
	}

	return b.String()
}
