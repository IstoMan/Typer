package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	sentence     string
	input        string
	correctStyle lipgloss.Style
	normalStyle  lipgloss.Style
	wrongStyle   lipgloss.Style
}

func InitialModel() Model {
	ns := lipgloss.NewStyle().
		Align(lipgloss.Center)

	cs := lipgloss.NewStyle().
		Inherit(ns).
		Foreground(lipgloss.Color("#30fc03"))

	ws := lipgloss.NewStyle().Inherit(ns).
		Foreground(lipgloss.Color("#fc0324"))

	return Model{
		sentence:     "The quick brown fox jumps over the lazy dog.",
		input:        "",
		correctStyle: cs,
		normalStyle:  ns,
		wrongStyle:   ws,
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
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			if len(m.input) < len(m.sentence) {
				m.input += string(msg.Runes)
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var ui strings.Builder
	ui.WriteString(m.normalStyle.Render(m.sentence))

	for i, char := range m.input {
		if string(char) == string(m.sentence[i]) {
			ui.WriteString(m.correctStyle.Render(string(char)))
		} else {
			ui.WriteString(m.wrongStyle.Render(string(char)))
		}
	}

	return ui.String()
}

func main() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
