package tui

import (
	"slog/environment"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14
const defaultWidth = 20

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	// quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type Item string

func (i Item) FilterValue() string {
	return string(i)
}

// type model struct {
// 	list     list.Model
// 	choice   string
// 	quitting bool
// }

// func (m model) Init() tea.Cmd {
// 	return nil
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.WindowSizeMsg:
// 		m.list.SetWidth(msg.Width)
// 		return m, nil

// 	case tea.KeyMsg:
// 		switch keypress := msg.String(); keypress {
// 		case "ctrl+c":
// 			m.quitting = true
// 			return m, tea.Quit

// 		case "enter":
// 			i, ok := m.list.SelectedItem().(Item)
// 			if ok {
// 				m.choice = string(i)
// 			}
// 			return m, tea.Quit
// 		}
// 	}

// 	var cmd tea.Cmd
// 	m.list, cmd = m.list.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {

// 	if m.choice != "" {
// 		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
// 	}
// 	if m.quitting {
// 		return quitTextStyle.Render("Not hungry? That’s cool.")
// 	}
// 	return "\n" + m.list.View()
// }

func CreateProgram(envs *environment.Environments) *tea.Program {
	return tea.NewProgram(newModel(envs), tea.WithAltScreen())

}
