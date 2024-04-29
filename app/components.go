package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var hightlight = lipgloss.Style{}.Foreground(lipgloss.Color("7"))
var dim = lipgloss.Style{}.Foreground(lipgloss.Color("8"))
var green = lipgloss.Style{}.Foreground(lipgloss.Color("2"))
var blue = lipgloss.Style{}.Foreground(lipgloss.Color("4"))

func presets(m Model) string {
	// Check if any value from the presets match the current value
	p := m.findMatchingPreset()

	b := strings.Builder{}

	b.WriteString("Presets: ")

	for i, preset := range m.presets {

		if p == preset.Name {
			b.WriteString(green.Render(fmt.Sprintf("[%d] %s", i+1, preset.Name)))
		} else {
			b.WriteString(dim.Render(fmt.Sprintf("[%d] %s", i+1, preset.Name)))
		}

		if i < len(m.presets)-1 {
			b.WriteString(dim.Render(" | "))
		}

	}

	return lipgloss.Style{}.MarginBottom(1).MarginTop(1).Render(b.String())
}

func truncateString(str string, maxLength int) string {
	if len(str) <= maxLength {
		return str
	}
	return fmt.Sprintf("%s..", str[:maxLength-2])
}

func choices(m Model) string {
	var s []string

	for i, choice := range m.variables {
		cursor := "  "
		if m.cursor == i {
			cursor = blue.Render("→ ")
		}

		cs := lipgloss.Style{}.
			Width(8).
			AlignHorizontal(lipgloss.Center).
			Render(truncateString(choice.Values[choice.Selected], 8))

		c := fmt.Sprintf("[%s]", cs)

		if m.cursor == i {
			c = blue.Render(c)
		}

		n := hightlight.Copy().MarginLeft(2).Width(8).Render(choice.Name)
		if m.cursor == i {
			n = blue.Copy().MarginLeft(2).Width(8).Render(choice.Name)
		}

		var d string
		if m.cursor == i {
			d += dim.Copy().MarginLeft(2).Render(choice.Description)
		}

		s = append(s, cursor+c+n+d)
	}

	s = append(s, "")

	return lipgloss.JoinVertical(lipgloss.Left, s...)
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
