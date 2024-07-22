package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/smueller264/converter/cli"
)

func main() {
	//Setting up logging to logfile
	file, err := os.OpenFile("errorlog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	//Starting the cli
	p := tea.NewProgram(cli.M, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Theres has been an error: %v", err)
		os.Exit(1)
	}

}
