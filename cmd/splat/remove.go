// remove.go
package main

import (
	"gopkg.in/urfave/cli.v2"
)

func removeInit() {
	CommandOpts = append(CommandOpts, &cli.Command{
		Name:    "remove",
		Aliases: []string{"R", "rm"},
		Usage:   "Remove a given package",
		Action:  Stub,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "force",
				Usage:   "Force removal of package even if it breaks things",
				Value:   Config.Remove.Force,
				Aliases: []string{"f"},
			},
		},
	})
}

func init() {
	Handlers = append(Handlers, removeInit)
}
