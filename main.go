package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	sentence     string
	input        string
	correctStyle lipgloss.Style
	normalStyle  lipgloss.Style
}

func InitialModel() Model {
	cs := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("#30fc03"))

	ns := lipgloss.NewStyle().
		Align(lipgloss.Center)

	return Model{
		sentence:     "The quick brown box jumps over the lazy",
		input:        "",
		correctStyle: cs,
		normalStyle:  ns,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			m.input = m.input[:len(m.input)-1]
		default:
			m.input += string(msg.Runes)
		}
	}

	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("%s%s\n\n%s", m.normalStyle.Render(m.sentence), m.correctStyle.Render(m.input), "ctrl+c to exit")
}

func main() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
