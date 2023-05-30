package tui

import (
	"slog/environment"

	"github.com/charmbracelet/bubbles/list"
)

func GetKs(envs environment.Environments) []list.Item {
	items := []list.Item{}
	for k, _ := range envs.Kubernetes {
		items = append(items, Item(k))
	}

	return items
}

// func GetDs(envs *environment.Environments) []list.Item {
// 	items := []list.Item{}
// 	for _, s := range envs.Kubernetes {
// 		items = append(items, Item(s.Name))
// 	}

// 	return items
// }
