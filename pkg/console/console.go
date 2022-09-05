package console

import (
	"fmt"
	"github.com/mgutz/ansi"
	"os"
)

func Success(message string) {
	colorOut(message, "green")
}

func Error(message string) {
	colorOut(message, "red")
}

func Warning(message string) {
	colorOut(message, "yellow")
}

func Exit(message string) {
	Error(message)
	os.Exit(1)
}

func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

func colorOut(message, color string) {
	fmt.Fprintln(os.stdout, ansi.Color(message, color))
}
