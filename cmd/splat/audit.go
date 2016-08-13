// audit.go
package main

import (
	"gopkg.in/urfave/cli.v2"
)

func auditInit() {
	CommandOpts = append(CommandOpts, &cli.Command{
		Name:    "audit",
		Aliases: []string{"A", "au"},
		Usage:   "Audit installed packages vs filesystem",
		Action:  Stub,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "report",
				Usage:   "Report on differences only.",
				Value:   Config.Audit.Report,
				Aliases: []string{"p"},
			},
			&cli.BoolFlag{
				Name:    "revert",
				Usage:   "Revert any filesystem changes to match recorded state",
				Value:   Config.Audit.Revert,
				Aliases: []string{"v"},
			},
		},
	})
}

func init() {
	Handlers = append(Handlers, auditInit)
}
