package flag

import (
	"gopkg.in/urfave/cli.v1"
)

var UpdateFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "force",
		Usage: "fuerza recrear el projecto",
	},
}

var InitFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "name",
		Usage: "nombre del projecto",
		Value: "example-project",
	},
	cli.StringFlag{
		Name:  "desc",
		Usage: "descripción del projecto",
		Value: "example description",
	},
	cli.StringFlag{
		Name:  "purpose",
		Usage: "propósito del projecto",
		Value: "example purpose",
	},
	cli.StringFlag{
		Name:  "language",
		Usage: "lenguaje de programacion",
		Value: "Go",
	},
	cli.StringFlag{
		Name:  "framework",
		Usage: "marco de trabajo",
	},
	cli.StringFlag{
		Name:  "orgunit",
		Usage: "unidad organizacional",
		Value: "Interno",
	},
	cli.StringSliceFlag{
		Name:  "deps",
		Usage: "dependencias de instalación",
	},
	cli.StringSliceFlag{
		Name:  "builddeps",
		Usage: "dependencias de compilación",
		Value: &cli.StringSlice{"golang"},
	},
}
