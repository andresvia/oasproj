package main

import (
	"github.com/andresvia/oasproj/command"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	for _, commonFlag := range commonFlags {
		updateFlags = append(updateFlags, commonFlag)
		initFlags = append(initFlags, commonFlag)
	}
	app := cli.NewApp()
	app.Name = "oasproj"
	app.Usage = "Maneja proyectos de la OAS"
	app.Version = "1.0.1"
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "show",
			Usage:  "muestra la información del proyecto",
			Action: command.Show,
		},
		cli.Command{
			Name:   "init",
			Usage:  "inicializa un projecto",
			Action: command.Init,
			Flags:  initFlags,
		},
		cli.Command{
			Name:   "update",
			Usage:  "actualiza un projecto",
			Action: command.Update,
			Flags:  updateFlags,
		},
		cli.Command{
			Name:   "check",
			Usage:  "verifica el proyecto",
			Action: command.Check,
		},
		cli.Command{
			Name:   "builddeps",
			Usage:  "instala dependencias de compilación",
			Action: command.Builddeps,
		},
		//cli.Command{
		//	Name:   "validate",
		//	Usage:  "valida que el proyecto este listo para publicación",
		//	Action: command.Validate,
		//},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var commonFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "with-daemon",
		Usage: "Crea el archivo de inicialización de systemd",
	},
}

var updateFlags = []cli.Flag{
	cli.BoolFlag{
		Name:   "force",
		Usage:  "fuerza recrear el projecto",
		EnvVar: "OAS_PROJ_FORCE_CREATE",
	},
}

var initFlags = []cli.Flag{
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
