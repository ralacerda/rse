package app

import (
	"github.com/charmbracelet/lipgloss"
)

func buildUi(m Model, f ...func(Model) string) string {

	var r []string

	for _, fn := range f {
		r = append(r, fn(m))
	}

	s := lipgloss.JoinVertical(lipgloss.Left, r...)

	return s
}
