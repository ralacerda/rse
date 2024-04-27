package main

import (
	"fmt"
	"os"

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

type model struct {
	cursor    int
	variables []variable
	presets   []preset
}

func applyPreset(m model, presetIndex int) model {
	for i, choice := range m.variables {
		if val, ok := m.presets[presetIndex].values[choice.name]; ok {
			m.variables[i].selected = val
		} else {
			// Reset the value to default if not on the preset
			m.variables[i].selected = 0
		}
	}
	return m
}

func initialModel() model {
	return model{
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

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
		case "1":
			m = applyPreset(m, 0)
		case "2":
			m = applyPreset(m, 1)
		case "3":
			m = applyPreset(m, 2)
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	var p string

	// Check if any value from the presets match the current value
	for _, preset := range m.presets {
		match := true
		for _, choice := range m.variables {
			if preset.values[choice.name] != choice.selected {
				match = false
				break
			}
		}

		if match {
			p = preset.name
			break
		}
	}

	s += "Presets: \n"
	for i, preset := range m.presets {

		if p == preset.name {
			s += "✓"
		}

		s += fmt.Sprintf("[%d] %s | ", i+1, preset.name)

	}
	s += "\n\n"

	for i, choice := range m.variables {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		value := choice.values[choice.selected]

		s += fmt.Sprintf("%s [%s] %s // %s \n", cursor, value, choice.name, choice.description)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	r, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if m, ok := r.(model); ok {
		for _, choice := range m.variables {
			fmt.Printf("%s: %v\n", choice.name, choice.values[choice.selected])
		}
	}
}
