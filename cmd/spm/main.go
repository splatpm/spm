package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli.v2"
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
		Name:    "spm",
		Usage:   "Splat Package Manager",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "Turn on debugging output",
			},
			&cli.BoolFlag{
				Name:  "nocolor",
				Value: false,
				Usage: "Turn off colors in program output",
			},
		},
		Commands: []*cli.Command{
			/*
				The audit command
			*/
			{
				Name:    "audit",
				Aliases: []string{"A"},
				Usage:   "Audit installed packages vs filesystem",
				Action:  Stub,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "report",
						Usage: "Report on differences only.",
						Value: true,
					},
					&cli.BoolFlag{
						Name:  "revert",
						Usage: "Revert any filesystem changes to match recorded state",
						Value: false,
					},
				},
			},
			/*
				The install command
			*/
			{
				Name:    "install",
				Aliases: []string{"I"},
				Usage:   "Install a given package",
				Action:  Stub,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "root",
						Usage: "Specify an alternate installation root",
						Value: "/",
					},
				},
			},
			/*
				The remove command
			*/
			{
				Name:    "remove",
				Aliases: []string{"R"},
				Usage:   "Remove a given package",
				Action:  Stub,
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}
	app.Run(os.Args)
	return
}
