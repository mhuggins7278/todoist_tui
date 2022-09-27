package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)
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

