package main

import (
	"github.com/andresvia/oasproj/command"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "oasproj"
	app.Usage = "Maneja proyectos de la OAS"
	app.Version = "1.0.0"
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
		},
		cli.Command{
			Name:   "update",
			Usage:  "actualiza un projecto",
			Action: command.Update,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:   "force",
					Usage:  "fuerza recrear el projecto",
					EnvVar: "OAS_PROJ_FORCE_CREATE",
				},
			},
		},
		cli.Command{
			Name:   "check",
			Usage:  "verifica el proyecto",
			Action: command.Check,
		},
		cli.Command{
			Name:   "validate",
			Usage:  "valida que el proyecto este listo para publicación",
			Action: command.Validate,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
