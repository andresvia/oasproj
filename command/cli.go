package command

import (
	"github.com/andresvia/oasproj/flag"
	"gopkg.in/urfave/cli.v1"
)

var Commands = []cli.Command{
	cli.Command{
		Name:   "show",
		Usage:  "muestra la información del proyecto",
		Action: Show,
	},
	cli.Command{
		Name:   "init",
		Usage:  "inicializa un proyecto",
		Action: Init,
		Flags:  flag.InitFlags,
	},
	cli.Command{
		Name:   "update",
		Usage:  "actualiza un proyecto",
		Action: Update,
		Flags:  flag.UpdateFlags,
	},
	cli.Command{
		Name:   "check",
		Usage:  "verifica el proyecto",
		Action: Check,
	},
	cli.Command{
		Name:   "builddeps",
		Usage:  "instala dependencias de compilación",
		Action: Builddeps,
	},
}
