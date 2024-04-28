package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ralacerda/rse/app"
)

func main() {
	p := tea.NewProgram(app.InitialModel())

	r, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if m, ok := r.(app.Model); ok {
		fmt.Print("\n")
		m.Output()
	}
}
