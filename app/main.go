package app

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type Variable struct {
	Name        string
	Description string
	Selected    int
	Values      []string
}

type Preset struct {
	Name   string
	Values map[string]int
}

type Model struct {
	cursor    int
	variables []Variable
	presets   []Preset
}

func (p Preset) apply(m Model) Model {
	for i, choice := range m.variables {
		if val, ok := p.Values[choice.Name]; ok {
			m.variables[i].Selected = val
		} else {
			// Reset the value to default if not on the preset
			m.variables[i].Selected = 0
		}
	}
	return m
}

func (p Preset) match(m Model) bool {
	for _, choice := range m.variables {
		if p.Values[choice.Name] != choice.Selected {
			return false
		}
	}
	return true
}

func (m Model) findMatchingPreset() string {
	for _, preset := range m.presets {
		if preset.match(m) {
			return preset.Name
		}
	}
	return ""
}

func New(v []Variable, p []Preset) Model {
	return Model{
		variables: v,
		presets:   p,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j", "tab":
			if m.cursor < len(m.variables)-1 {
				m.cursor++
			}
		case " ", "l", "right":
			if m.variables[m.cursor].Selected+1 < len(m.variables[m.cursor].Values) {
				m.variables[m.cursor].Selected++
			} else {
				m.variables[m.cursor].Selected = 0
			}
		case "h", "left":
			if m.variables[m.cursor].Selected > 0 {
				m.variables[m.cursor].Selected--
			} else {
				m.variables[m.cursor].Selected = len(m.variables[m.cursor].Values) - 1
			}
		case "enter":

		default:
			if num, err := strconv.Atoi(msg.String()); err == nil && num >= 1 && num <= len(m.presets) {
				m = m.presets[num-1].apply(m)
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	return buildUi(m, presets, choices, footer)
}

func (m Model) Output() map[string]string {
	envs := make(map[string]string)
	for _, choice := range m.variables {
		envs[choice.Name] = choice.Values[choice.Selected]
	}
	return envs
}
