package command

import (
	"errors"
	"fmt"
	"github.com/andresvia/oasproj/cliutil"
	"github.com/andresvia/oasproj/project"
	"github.com/andresvia/oasproj/template"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/exec"
	"path/filepath"
)

func projectGit(ctx *cli.Context) string {
	return filepath.Join(cliutil.ProjectHome(ctx), ".git")
}

func projectHomeExists(ctx *cli.Context) bool {
	if _, err := os.Stat(cliutil.ProjectHome(ctx)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func projectGitExists(ctx *cli.Context) bool {
	if _, err := os.Stat(projectGit(ctx)); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func ensureProjectHome(ctx *cli.Context) (err error) {
	if !projectHomeExists(ctx) {
		if err = os.MkdirAll(cliutil.ProjectHome(ctx), 0755); err == nil {
			err = gitInit(ctx)
		}
	} else if !projectGitExists(ctx) {
		err = gitInit(ctx)
	}
	return
}

func ensureProjectGit(ctx *cli.Context) (err error) {
	if !projectGitExists(ctx) {
		err = gitInit(ctx)
	}
	return
}

func gitInit(ctx *cli.Context) (err error) {
	var out []byte
	pwd, _ := os.Getwd()
	if err = os.Chdir(cliutil.ProjectHome(ctx)); err == nil {
		git_init := exec.Command("git", "init")
		if out, err = git_init.CombinedOutput(); err == nil {
			fmt.Printf("%s", string(out[:]))
		}
	}
	os.Chdir(pwd)
	return
}

func createOrUpdateProjectFiles(ctx *cli.Context) (err error) {
	this_project := project.LoadProject(cliutil.ProjectHome(ctx))
	if errs := template.AllTemplates(ctx, this_project); len(errs) > 0 {
		err = errs[0]
	}
	return
}

func initializeProjectDescriptorFile(ctx *cli.Context) (err error) {
	this_project := project.New(ctx)
	if err = ensureProjectHome(ctx); err != nil {
		return
	}
	if err = ensureProjectGit(ctx); err != nil {
		return
	}
	err = this_project.WriteFile(cliutil.ProjectHome(ctx))
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

func Validate(ctx *cli.Context) (err error) {
	// test pasan
	return
}
