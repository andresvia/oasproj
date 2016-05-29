package command

import (
	"errors"
	"fmt"
	"github.com/andresvia/oasproj/project"
	"github.com/andresvia/oasproj/template"
	"gopkg.in/urfave/cli.v1"
)

func projectHome(ctx *cli.Context) (home string) {
	home = ctx.Args().First()
	return
}

func createOrUpdateProjectFiles(ctx *cli.Context) (err error) {
	this_project := project.LoadProject(projectHome(ctx))
	if errs := template.AllTemplates(projectHome(ctx), this_project, ctx.Bool("force")); len(errs) > 0 {
		err = errs[0]
	}
	return
}

var noProjectError = errors.New("No se ha inicializado el proyecto, inicializar con 'init'")

func Show(ctx *cli.Context) (err error) {
	if project.MetadataExists(projectHome(ctx)) {
		this_project := project.LoadProject(projectHome(ctx))
		fmt.Println(this_project)
	} else {
		err = noProjectError
	}
	return
}

func Check(ctx *cli.Context) (err error) {
	if !project.MetadataExists(projectHome(ctx)) {
		err = noProjectError
	}
	return
}

func Init(ctx *cli.Context) (err error) {
	if !project.MetadataExists(projectHome(ctx)) {
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
