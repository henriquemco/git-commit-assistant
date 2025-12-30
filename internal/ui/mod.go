package ui

import (
	"fmt"
	"time"

	lg "github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

var StyleError = lg.NewStyle().Foreground(lg.Color("#D8647E")).Bold(true).Render

var Bold = lg.NewStyle().Bold(true).Render

var StyleCommit = lg.NewStyle().
	Bold(true).
	Foreground(lg.Color("#8AA46E")).Render

func Introduction() {
	banner := []string{
		"  ____ ___ _____      _    ____ ____ ___ ____ _____  _    _   _ _____",
		" / ___|_ _|_   _|    / \\  / ___/ ___|_ _/ ___|_   _|/ \\  | \\ | |_   _|",
		"| |  _ | |  | |     / _ \\ \\___ \\___ \\| |\\___ \\ | | / _ \\ |  \\| | | |",
		"| |_| || |  | |    / ___ \\ ___) |__) | | ___) || |/ ___ \\| |\\  | | |  ",
		" \\____|___| |_|   /_/   \\_\\____/____/___|____/ |_/_/   \\_\\_| \\_| |_| ",
		"                I'm an AI-powered Git commit assistant.                 ",
	}

	for _, line := range banner {
		fmt.Println(line)
	}
}

func Loading(stopchan chan struct{}) {
	frames := []string{"⣾ ", "⣽ ", "⣻ ", "⢿ ", "⡿ ", "⣟ ", "⣯ ", "⣷ "}
	i := 0
	for {
		select {
		case <-stopchan:

			fmt.Print("\r                          \r")
			fmt.Print("\n")
			return
		default:
			fmt.Println(Bold(fmt.Sprintf("\r    %s Generating... \r ", frames[i%len(frames)])))
			time.Sleep(200 * time.Millisecond)
			i++
		}
	}
}
