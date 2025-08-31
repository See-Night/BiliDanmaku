package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DanmakuModel struct {
	username string
	context  string

	width int
}

func (d DanmakuModel) Init() tea.Cmd {
	return nil
}

func (d DanmakuModel) Update(msg tea.Msg) (DanmakuModel, tea.Cmd) {
	return d, nil
}

func (d DanmakuModel) View() string {
	nameStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
	systemStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("160")).Bold(true)
	if d.username == "NOTICE" {
		return systemStyle.Render(fmt.Sprintf("%s", d.context)) + "\n"
	}
	return fmt.Sprintf("%s %s\n", nameStyle.Render(d.username), d.context)
}

func NewDanmaku(username string, context string) DanmakuModel {
	return DanmakuModel{
		username: username,
		context:  context,
	}
}
