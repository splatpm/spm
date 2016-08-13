// install.go
package main

import (
	"gopkg.in/urfave/cli.v2"
)

func init() {
	CommandOpts = append(CommandOpts, &cli.Command{
		Name:    "install",
		Aliases: []string{"I", "in"},
		Usage:   "Install a given package",
		Action:  Stub,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "root",
				Usage:   "Specify an alternate installation root",
				Value:   "/",
				Aliases: []string{"r"},
			},
		},
	})
}
