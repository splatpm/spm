package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

var (
	CommandOpts []*cli.Command
)

func Stub(c *cli.Context) error {
	return nil
}

func process() {
	fmt.Println("process()")
	return
}

func main() {
	app := &cli.App{
		Name:    "splat",
		Usage:   "Splat Package Manager",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Value:   false,
				Usage:   "Turn on debugging output",
				Aliases: []string{"D"},
			},
			&cli.BoolFlag{
				Name:    "nocolor",
				Value:   false,
				Usage:   "Turn off colors in program output",
				Aliases: []string{"C"},
			},
		},
		Commands: CommandOpts,
		Action: func(c *cli.Context) error {
			return nil
		},
	}
	app.Run(os.Args)
	return
}
