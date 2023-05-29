package tui

import (
	"fmt"
	"io"
	"slog/environment"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	tui_env = iota
	tui_k8s_cluster
	tui_grepper
	Kubernetes = "Kubernetes"
	Docker     = "Docker"
)

type model struct {
	environmentList     list.Model
	environmentSelected string // docker or k8s
	clusterList         list.Model
	clusterSelected     string
	cursor              int
	namespace           string
	selectedView        int
	filter              string // grep filter to apply
}

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func createEnvironmentList() []list.Item {
	items := []list.Item{}
	items = append(items, Item(Kubernetes))
	items = append(items, Item(Docker))

	return items
}

func newModel(envs *environment.Environments) model {
	envList := list.New(createEnvironmentList(), itemDelegate{}, defaultWidth, listHeight)
	envList.Title = "Select an environment"
	envList.SetShowStatusBar(false)
	envList.SetFilteringEnabled(true)
	envList.Styles.Title = titleStyle
	envList.Styles.PaginationStyle = paginationStyle
	envList.Styles.HelpStyle = helpStyle

	// k8sClusters := GetKs(envs)
	k8sClusterList := list.New(GetKs(envs), itemDelegate{}, defaultWidth, listHeight)
	k8sClusterList.Title = "Select a cluster to use"
	k8sClusterList.SetShowStatusBar(false)
	k8sClusterList.SetFilteringEnabled(true)
	k8sClusterList.Styles.Title = titleStyle
	k8sClusterList.Styles.PaginationStyle = paginationStyle
	k8sClusterList.Styles.HelpStyle = helpStyle

	m := model{
		environmentList:     envList,
		environmentSelected: "",
		clusterList:         k8sClusterList,
	}

	return m
}

func (m model) View() string {
	if m.selectedView == tui_env || m.environmentSelected == "" {
		return m.environmentList.View()
	}

	if m.selectedView == tui_k8s_cluster {
		return m.clusterList.View()
	}

	fmt.Printf("initial view!")
	// This is never rendered
	return m.environmentList.View()
}

func (m model) Init() tea.Cmd {
	// tea.WithAltScreen()
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Global events
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.environmentList.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			// m.quitting = true
			return m, tea.Quit
		}
	}

	// TUI environment view
	if m.selectedView == tui_env {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
			case "enter":
				i, ok := m.environmentList.SelectedItem().(Item)
				if ok {
					m.environmentSelected = i.FilterValue()
					m.selectedView = tui_k8s_cluster
				}
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.environmentList, cmd = m.environmentList.Update(msg)
		return m, cmd
	}

	// TUI kubernetes environment view
	if m.selectedView == tui_k8s_cluster {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch keypress := msg.String(); keypress {
			case "enter":
				i, ok := m.clusterList.SelectedItem().(Item)
				if ok {
					m.clusterSelected = i.FilterValue()
					m.selectedView = tui_k8s_cluster
				}
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.clusterList, cmd = m.clusterList.Update(msg)
		return m, cmd
	}

	return m, nil
}
