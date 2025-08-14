package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	sentence   string
	typedText  string
	startTime  time.Time
	isOver     bool
	hasStarted bool
	wpm        float64
}

var (
	normalStyle = lipgloss.NewStyle().
			Align(lipgloss.Center)

	correctStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#30fc03"))

	wrongStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#fc0324"))

	helpStyle = lipgloss.NewStyle().
			Align(lipgloss.Bottom).
			Italic(true).
			Faint(true)
)

func getQuote() string {
	file, err := os.Open("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}

	randomNumber := 0 + rand.Intn((len(quotes)-1)-0+1)
	return quotes[randomNumber]
}

func InitialModel() Model {
	return Model{
		sentence:   getQuote(),
		typedText:  "",
		isOver:     false,
		hasStarted: false,
		wpm:        0,
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
			if len(m.typedText) > 0 {
				m.typedText = m.typedText[:len(m.typedText)-1]
			}
		case "alt+backspace":
			if len(m.typedText) > 0 {
				for i := len(m.typedText) - 1; i > 0; i-- {
					if string(m.typedText[i]) == " " {
						m.typedText = m.typedText[:i+1]
						break
					}
				}
			}
		default:
			if len(m.typedText) < len(m.sentence) {
				if !m.hasStarted {
					m.startTime = time.Now()
					m.hasStarted = true
				}

				m.typedText += string(msg.Runes)

				if m.sentence == m.typedText {
					m.isOver = true
					elapsed := time.Since(m.startTime).Minutes()

					words := float64(len(m.sentence)) / 5
					m.wpm = words / elapsed
				}
			}
		}
	}

	return m, nil
}

func (m Model) View() string {
	var ui strings.Builder

	if m.isOver {
		congrats := lipgloss.NewStyle().Foreground(lipgloss.Color("#30fc03")).Align(lipgloss.Center).Border(lipgloss.RoundedBorder(), true)
		s := fmt.Sprintf("Your final wpm is: %1.f", m.wpm)
		ui.WriteString(congrats.Render(s))
		ui.WriteString("\n\n")
		ui.WriteString(helpStyle.Render("ctrl+c to exit"))
		return ui.String()
	}

	for i, char := range m.sentence {
		if i < len(m.typedText) {
			if m.typedText[i] == byte(char) {
				ui.WriteString(correctStyle.Render(string(char)))
			} else {
				ui.WriteString(wrongStyle.Render(string(char)))
			}
		} else {
			ui.WriteString(normalStyle.Render(string(char)))
		}
	}

	ui.WriteString("\n\n")

	ui.WriteString(helpStyle.Render("ctrl+c to exit"))

	return ui.String()
}

func main() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
