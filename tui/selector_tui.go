package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type selectorModel struct {
	title   string
	choices []string
	cursor  int
}

func (m *selectorModel) up() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *selectorModel) down() {
	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
}

func (m selectorModel) Result() int {
	return m.cursor
}

func initialSelectorModel(title string, choices []string) selectorModel {
	return selectorModel{
		title:   title,
		choices: choices,
	}
}

func (m selectorModel) Init() tea.Cmd {
	return nil
}

func (m selectorModel) View() string {
	s := fmt.Sprintf("%s\n\n", m.title)

	for ii, choice := range m.choices {
		cursor := " "
		if m.cursor == ii {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}

func (m selectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.up()

		case "down", "j":
			m.down()

		case "enter", " ":
			return m, tea.Quit
		}
	}

	return m, nil
}

func RunTuiSelector(title string, choices []string) (int, error) {
	p := tea.NewProgram(initialSelectorModel(title, choices))
	model, err := p.StartM()
	if err != nil {
		return 0, err
	}

	return model.(selectorModel).Result(), nil
}
