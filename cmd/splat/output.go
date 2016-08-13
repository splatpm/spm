// output.go
package main

import (
	"fmt"
	"os"

	"github.com/splatpm/gout"
)

func output(tag string, msg string) {
	fmt.Printf("%s %s\n", tag, msg)
}

func Debug(msg string) {
	if Config.Debug {
		if Config.Color {
			output(fmt.Sprintf("%s%s",
				gout.String("DEBUG").Cyan(),
				gout.String(":").Bold().White()),
				msg)
		} else {
			output("DEBUG:", msg)
		}
	}
}

func Error(msg string) {
	tag := ""
	if Config.Color {
		tag = fmt.Sprintf("%s%s",
			gout.String("ERROR").Bold().Red(),
			gout.String(":").Bold().White())
	} else {
		tag = "ERROR:"
	}
	output(tag, msg)
	os.Exit(1)
}
