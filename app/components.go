package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var hightlight = lipgloss.Style{}.Foreground(lipgloss.Color("7"))
var dim = lipgloss.Style{}.Foreground(lipgloss.Color("8"))

func header(_ Model) string {
	s := hightlight
	return s.Render("What should we buy?")
}

func footer(_ Model) string {

	k := []string{
		"q: quit",
		"↑/↓: navigate",
		"←/→: change value",
		"1-9: select preset",
		"enter: submit",
		"?: help",
	}

	t := ""

	for i, v := range k {
		// Split the key and the command
		s := strings.Split(v, ":")
		t += hightlight.Render(s[0]) + dim.Render(s[1])
		if i < len(k)-1 {
			t += dim.Render(" | ")
		}
	}

	t += "\n"

	return t
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
