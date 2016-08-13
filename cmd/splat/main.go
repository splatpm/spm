package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"
)

var (
	CommandOpts []*cli.Command
	Config      *SplatConfig
	Handlers    []func()
)

func Stub(c *cli.Context) error {
	return nil
}

func init() {
	BuildConfig()
}

func main() {
	for _, v := range Handlers {
		v()
	}

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
			// Every "Action" should do this for it's arguments
			// at the beginning of execution time to ensure all
			// desired command line options are set.
			if c.Bool("debug") && !Config.Debug {
				Config.Debug = true
			}
			if c.Bool("color") && !Config.Color {
				Config.Color = true
			}
			Debug(fmt.Sprintf("%+v", Config))
			return nil
		},
	}
	app.Run(os.Args)
	return
}
