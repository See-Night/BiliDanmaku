package components

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SpinnerModel struct {
	spinner spinner.Model
	context string
	err     error
}

func (s SpinnerModel) Init() tea.Cmd {
	return s.spinner.Tick
}

func (c SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case error:
		c.err = msg
		return c, nil
	default:
		var cmd tea.Cmd
		c.spinner, cmd = c.spinner.Update(msg)
		return c, cmd
	}
}

func (c SpinnerModel) View() string {
	if c.err != nil {
		return "\n\n" + c.err.Error() + "\n\n"
	}
	str := "\n" + c.spinner.View() + c.context + "\n\n"
	return str
}

func (c *SpinnerModel) SetErr(err error) {
	c.err = err
}

func NewSpinnerModel(context string) SpinnerModel {
	s := SpinnerModel{
		spinner: spinner.New(),
		context: context,
	}
	s.spinner.Spinner = spinner.Dot
	s.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return s
}
