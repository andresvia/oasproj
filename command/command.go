package command

import (
	"errors"
	"fmt"
	"github.com/andresvia/oasproj/cliutil"
	"github.com/andresvia/oasproj/errutil"
	"github.com/andresvia/oasproj/project"
	"github.com/andresvia/oasproj/template"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/exec"
	"strings"
)

func ensureProjectHome(ctx *cli.Context) (err error) {
	if err = os.MkdirAll(cliutil.ProjectHome(ctx), 0755); err == nil {
		err = ensureProjectGit(ctx)
	}
	return
}

func ensureProjectGit(ctx *cli.Context) (err error) {
	var out []byte
	pwd, _ := os.Getwd()
	if err = os.Chdir(cliutil.ProjectHome(ctx)); err == nil {
		git_init := exec.Command("git", "init")
		if out, _ = git_init.CombinedOutput(); err == nil {
			fmt.Printf("%s", string(out[:]))
		}
	}
	os.Chdir(pwd)
	return
}

func createOrUpdateProjectFiles(ctx *cli.Context) (err error) {
	errs := template.DoTemplates(ctx)
	err = errutil.FirstOrNil(errs)
	return
}

func initializeProjectDescriptorFile(ctx *cli.Context) (err error) {
	this_project := project.New(ctx)
	if err = ensureProjectHome(ctx); err == nil {
		err = this_project.WriteFile(cliutil.ProjectHome(ctx))
	}
	return
}

var noProjectError = errors.New("No se ha inicializado el proyecto, inicializar con 'init'")

func Show(ctx *cli.Context) (err error) {
	if project.MetadataExists(cliutil.ProjectHome(ctx)) {
		this_project := project.LoadProject(cliutil.ProjectHome(ctx))
		fmt.Println(this_project)
	} else {
		err = noProjectError
	}
	return
}

func Check(ctx *cli.Context) (err error) {
	if !project.MetadataExists(cliutil.ProjectHome(ctx)) {
		err = noProjectError
	}
	return
}

func Init(ctx *cli.Context) (err error) {
	if !project.MetadataExists(cliutil.ProjectHome(ctx)) {
		if err = initializeProjectDescriptorFile(ctx); err == nil {
			err = createOrUpdateProjectFiles(ctx)
		}
	} else {
		err = errors.New("Ya se ha inicializado el proyecto, actualizar con 'update'")
	}
	return
}

func Update(ctx *cli.Context) (err error) {
	err = createOrUpdateProjectFiles(ctx)
	return
}

func Builddeps(ctx *cli.Context) (err error) {
	if project.MetadataExists(cliutil.ProjectHome(ctx)) {
		this_project := project.LoadProject(cliutil.ProjectHome(ctx))
		install_args := []string{"install", "-y"}
		builddeps := this_project.Build_dependencies["os"]
		for _, builddep := range builddeps {
			if strings.IndexAny(builddep, "()") == -1 {
				install_args = append(install_args, builddep)
			}
		}
		if len(install_args) > 2 {
			install := exec.Command("yum", install_args...)
			if err = install.Start(); err == nil {
				fmt.Println("Esperando a la instalación de depedendencias de compilación")
				err = install.Wait()
			}
		}
	} else {
		err = noProjectError
	}
	return
}
