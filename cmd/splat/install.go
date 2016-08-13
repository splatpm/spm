// install.go
package main

import (
	"gopkg.in/urfave/cli.v2"
)

func installInit() {
	CommandOpts = append(CommandOpts, &cli.Command{
		Name:    "install",
		Aliases: []string{"I", "in"},
		Usage:   "Install a given package",
		Action:  Stub,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "root",
				Usage:   "Specify an alternate installation root",
				Value:   Config.Install.Root,
				Aliases: []string{"r"},
			},
		},
	})
}

func init() {
	Handlers = append(Handlers, installInit)
}
