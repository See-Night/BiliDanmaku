package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	blurStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7"))
	focusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("0")).
			Background(lipgloss.Color("40")).
			Bold(true)
)

type ButtonModel struct {
	context string
	focused bool
}

func (b ButtonModel) Init() tea.Cmd {
	return nil
}

func (b ButtonModel) Update(msg tea.Msg) (ButtonModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if b.focused {
				return b, nil
			}
		}
	}

	return b, nil
}

func (b ButtonModel) View() string {
	if b.focused {
		return focusStyle.Render(fmt.Sprintf("%s", b.context))
	} else {
		return blurStyle.Render(fmt.Sprintf("%s", b.context))
	}
}

func (b *ButtonModel) Focus() {
	b.focused = true
}

func (b *ButtonModel) Blur() {
	b.focused = false
}

func NewButtonModel(context string) ButtonModel {
	return ButtonModel{
		context: context,
	}
}
