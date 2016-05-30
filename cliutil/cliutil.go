package cliutil

import "gopkg.in/urfave/cli.v1"

func ProjectHome(ctx *cli.Context) string {
	home := ctx.Args().First()
	if home == "" {
		return "."
	} else {
		return home
	}
}
