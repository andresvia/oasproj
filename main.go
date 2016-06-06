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
	app.Version = "1.0"
	app.EnableBashCompletion = true
	app.Commands = command.Commands
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
