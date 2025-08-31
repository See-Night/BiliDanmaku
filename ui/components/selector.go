package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	selectFocusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

type SelectorModel struct {
	options []string
	index   int
}

func (s SelectorModel) Init() tea.Cmd {
	return nil
}

func (s SelectorModel) Update(msg tea.Msg) (SelectorModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			s.index = (s.index - 1 + len(s.options)) % len(s.options)
		case "down":
			s.index = (s.index + 1) % len(s.options)
		}
	}

	return s, nil
}

func (s SelectorModel) View() string {
	var view strings.Builder
	for index, option := range s.options {
		if index == s.index {
			view.WriteString(selectFocusStyle.Render(fmt.Sprintf("> %s", option)))
		} else {
			view.WriteString(fmt.Sprintf("  %s", option))
		}
	}
	return view.String()
}

func (s *SelectorModel) GetValue() string {
	return s.options[s.index]
}

func NewSelectorModel(options []string) *SelectorModel {
	return &SelectorModel{
		options: options,
		index:   0,
	}
}
