package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ralacerda/rse/app"
)

func main() {
	a := app.New([]app.Variable{
		{
			Name:        "EDITOR",
			Description: "The editor to use",
			Selected:    0,
			Values:      []string{"vim", "emacs", "nano"},
		},
		{
			Name:        "SHELL",
			Description: "The shell to use",
			Selected:    0,
			Values:      []string{"bash", "zsh", "fish"},
		},
	}, []app.Preset{
		{
			Name: "Home",
			Values: map[string]int{
				"EDITOR": 1,
				"SHELL":  1,
			},
		},
		{
			Name: "Work",
			Values: map[string]int{
				"EDITOR": 0,
				"SHELL":  0,
			}}})

	p := tea.NewProgram(a)

	r, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	m, _ := r.(app.Model)
	envs := m.Output()

	cmd := exec.Command("printenv")

	for key, env := range envs {
		os.Setenv(key, env)
	}

	cmd.Stdout = os.Stdout // or any other io.Writer
	cmd.Stderr = os.Stderr // or any other io.Writer
	if err := cmd.Run(); err != nil {
		log.Println("ERROR Running")
	}

}
