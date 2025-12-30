package ui

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type modelSelect struct {
	cursor   int
	choices  []string
	selected map[int]struct{}
}

func Select() []string {
	p := tea.NewProgram(InitialModel())
	model, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	prompt := strings.Split(model.View(), "\n")
	return prompt
}

func InitialModel() modelSelect {
	return modelSelect{
		choices: []string{"Yes", "No"},

		selected: make(map[int]struct{}),
	}
}

func (m modelSelect) Init() tea.Cmd {
	return tea.SetWindowTitle("Grocery List")
}

func (m modelSelect) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":

			outher := ((m.cursor) - 1) * (-1)
			_, ok := m.selected[outher]
			if ok {
				delete(m.selected, outher)
			}

			_, ok = m.selected[m.cursor]

			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
			return m, tea.Quit

		}
	}

	return m, nil
}

func (m modelSelect) View() string {
	s := ""
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf(" %s [%s] %s\n", cursor, checked, choice)
	}

	return s
}
