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

func getProjects() (items []list.Item) {
	projects, err := get[Projects]("https://api.todoist.com/rest/v2/projects")
	if err != nil {
		return nil
	}
	//convert projest to list.Item interface
	for _, project := range projects {
		items = append(items, project)
	}
	return
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (m model) Init() tea.Cmd {
	return nil
}

func (p Project) FilterValue() string {
	return p.Name
}

func (p Project) Title() string {
	return p.Name
}

func (p Project) Description() string {
	return p.Name
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

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
