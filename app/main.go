package app

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type variable struct {
	name        string
	description string
	selected    int
	values      []string
}

type preset struct {
	name   string
	values map[string]int
}

type Model struct {
	cursor    int
	variables []variable
	presets   []preset
}

func (p preset) apply(m Model) Model {
	for i, choice := range m.variables {
		if val, ok := p.values[choice.name]; ok {
			m.variables[i].selected = val
		} else {
			// Reset the value to default if not on the preset
			m.variables[i].selected = 0
		}
	}
	return m
}

func (p preset) match(m Model) bool {
	for _, choice := range m.variables {
		if p.values[choice.name] != choice.selected {
			return false
		}
	}
	return true
}

func (m Model) findMatchingPreset() string {
	for _, preset := range m.presets {
		if preset.match(m) {
			return preset.name
		}
	}
	return ""
}

func InitialModel() Model {
	return Model{
		variables: []variable{
			{
				name:        "Food",
				description: "What should we buy at the market?",
				values:      []string{"Burguer", "Salad", "Vegan Burguer"},
			},
			{
				name:   "Drink",
				values: []string{"Water", "Coke", "Diet Coke"},
			},
		},
		presets: []preset{
			{
				name: "Default",
				values: map[string]int{
					"Food":  0,
					"Drink": 0,
				}},
			{
				name: "Vegan",
				values: map[string]int{
					"Food":  2,
					"Drink": 1,
				}},
			{
				name: "Healthy",
				values: map[string]int{
					"Food":  1,
					"Drink": 2,
				}},
		},
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
			if m.variables[m.cursor].selected+1 < len(m.variables[m.cursor].values) {
				m.variables[m.cursor].selected++
			} else {
				m.variables[m.cursor].selected = 0
			}
		case "h", "left":
			if m.variables[m.cursor].selected > 0 {
				m.variables[m.cursor].selected--
			} else {
				m.variables[m.cursor].selected = len(m.variables[m.cursor].values) - 1
			}

		default:
			if num, err := strconv.Atoi(msg.String()); err == nil && num >= 1 && num <= len(m.presets) {
				m = m.presets[num-1].apply(m)
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	return buildUi(m, header, presets, choices, footer)
}

func (m Model) Output() {
	for _, choice := range m.variables {
		fmt.Printf("%s: %v\n", choice.name, choice.values[choice.selected])
	}
}
