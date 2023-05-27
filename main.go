package main

import (
	"fmt"
	"os"
	"slog/file_parser"
	"slog/ui"

	"github.com/charmbracelet/bubbles/list"
)

func makeItems(e []file_parser.Environment) {

}

func main() {
	systems, err := file_parser.GetSystems()
	if err != nil {
		fmt.Printf("could not open the system environments file, %s \n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", systems)
	items := []list.Item{}

	for _, s := range systems {
		items = append(items, ui.Item(s.Name))
	}

	if _, err := ui.CreateProgram(items).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
