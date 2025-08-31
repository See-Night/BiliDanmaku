package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	labelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true)
)

type InputModel struct {
	input   textinput.Model
	label   string
	focused bool
	Width   int
	Height  int
}

func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		newWidth := msg.Width - 8
		valueLength := len(m.input.Value())
		if m.Width == 0 || valueLength < newWidth {
			m.input.Width = newWidth
			m.input.SetValue(m.input.Value())
		} else {
			m.input.Width = valueLength
		}
	case tea.KeyMsg:
		if m.Width != 0 && m.Width < len(m.input.Value()) {
			m.input.Width = len(m.input.Value()) + 1
			m.input.SetValue(m.input.Value())
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m InputModel) View() string {
	style := lipgloss.NewStyle()
	if m.Width != 0 {
		style = style.Width(m.Width).MaxHeight(4).Height(4)
	}
	return style.Render(labelStyle.Render(m.label) + "\n" + m.input.View())
}

func (m *InputModel) Focus() tea.Cmd {
	m.focused = true
	return m.input.Focus()
}

func (m *InputModel) Blur() {
	m.focused = false
	m.input.Blur()
}

func (m *InputModel) GetValue() string {
	return m.input.Value()
}

func (m *InputModel) SetValue(value string) {
	m.input.SetValue(value)
}

func NewInputModel(label string, placeholder string) InputModel {
	input := textinput.New()
	model := InputModel{
		input:   input,
		label:   label,
		focused: false,
		Width:   0,
	}
	model.input.Placeholder = placeholder
	model.input.Prompt = ""

	return model
}
