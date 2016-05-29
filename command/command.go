package command

import (
	"errors"
	"fmt"
	"github.com/andresvia/oasproj/project"
	"github.com/andresvia/oasproj/template"
	"gopkg.in/urfave/cli.v1"
)

func createOrUpdateProjectFiles(ctx *cli.Context) (err error) {
	this_project := project.LoadProject(ctx.Args().First())
	if errs := template.AllTemplates(ctx, this_project); len(errs) > 0 {
		err = errs[0]
	}
	return
}

var noProjectError = errors.New("No se ha inicializado el proyecto, inicializar con 'init'")

func Show(ctx *cli.Context) (err error) {
	if project.MetadataExists(ctx.Args().First()) {
		this_project := project.LoadProject(ctx.Args().First())
		fmt.Println(this_project)
	} else {
		err = noProjectError
	}
	return
}

func Check(ctx *cli.Context) (err error) {
	if !project.MetadataExists(ctx.Args().First()) {
		err = noProjectError
	}
	return
}

func Init(ctx *cli.Context) (err error) {
	if !project.MetadataExists(ctx.Args().First()) {
		err = createOrUpdateProjectFiles(ctx)
	} else {
		err = errors.New("Ya se ha inicializado el proyecto, actualizar con 'update'")
	}
	return
}

func Update(ctx *cli.Context) (err error) {
	err = createOrUpdateProjectFiles(ctx)
	return
}

func Validate(ctx *cli.Context) (err error) {
	// test pasan
	return
}
