package ui

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

func Loading(stopchan chan bool) {
	frames := []string{"|", "/", "-", "\\"}
	i := 0
	for {
		select {
		case <-stopchan:
			fmt.Print("\r                          \r")
			return
		default:
			color.RGB(192, 202, 245).Printf("\r%s", frames[i%len(frames)])
			time.Sleep(200 * time.Millisecond)
			i++
		}
	}
}
