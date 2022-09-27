package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// func main() {
// 	p := tea.NewProgram(initialModel())
// 	if err := p.Start(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}

// }


var docStyle = lipgloss.NewStyle().
  Margin(1, 2).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63"))


func main() {

	items := getProjects()

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Projects"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
