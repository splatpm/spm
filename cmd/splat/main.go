package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

var (
	CommandOpts []*cli.Command
	Config      *SplatConfig
)

func Stub(c *cli.Context) error {
	return nil
}

func main() {
	BuildConfig()
	app := &cli.App{
		Name:    "splat",
		Usage:   "Splat Package Manager",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Value:   Config.Debug,
				Usage:   "Turn on debugging output",
				Aliases: []string{"D"},
			},
			&cli.BoolFlag{
				Name:    "color",
				Value:   Config.Color,
				Usage:   "Turn on colors in program output",
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
