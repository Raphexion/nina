package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg error

type inputModel struct {
	title     string
	textInput textinput.Model
	err       error
}

func (m inputModel) Result() string {
	return m.textInput.Value()
}

func initialInputModel(title, text, placeholder string) inputModel {
	ti := textinput.NewModel()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 80

	if len(text) > 0 {
		ti.SetValue(text)
	}

	return inputModel{
		title:     title,
		textInput: ti,
		err:       nil,
	}
}

func (m inputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m inputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m inputModel) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n(esc to quit)\n", m.title, m.textInput.View())
}

func RunInput(title, text, placeholder string) (string, error) {
	p := tea.NewProgram(initialInputModel(title, text, placeholder))
	model, err := p.StartM()
	if err != nil {
		return "", err
	}

	return model.(inputModel).Result(), nil
}
