package main

import (
	"fmt"
	"himitsu/config"
	"himitsu/ui/components"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	config.Init()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type (
	tickMsg struct{}
)

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

type model struct {
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return tick()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			return m, tea.Quit
		}
	case tickMsg:
		return m, tick()
	}
	return m, cmd
}

func (m model) View() string {
	tpl := components.VersionNumber() + "\n\n"
	for _, totp := range config.Config {
		tpl += components.TOTP(totp)
	}
	tpl += "\n" + components.KeybindsHints([]string{"esc: quit"}) + "\n"
	return fmt.Sprintf(tpl)
}
