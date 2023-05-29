package main

import (
	"fmt"
	"os"
	"slog/environment"
	"slog/tui"
)

func main() {
	envs, err := environment.GetEnvironments()
	if err != nil {
		fmt.Printf("could not open the system environments file, %s \n", err)
		os.Exit(1)
	}

	if _, err := tui.CreateProgram(envs).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
