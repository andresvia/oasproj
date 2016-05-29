package template

import (
	"fmt"
	"github.com/andresvia/oasproj/project"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

type TemplateInfo struct {
	Project     project.Project
	Create      string
	Update      string
	ForceUpdate string
}

func CreateDefaultTemplateInfo(this_project project.Project) (template_info TemplateInfo) {
	template_info.Project = this_project
	template_info.Create = `Este archivo fue creado con "oasproj init"`
	template_info.Update = `Este archivo fue creado con "oasproj init" y será sobre-escrito con "oasproj update"`
	template_info.ForceUpdate = `Este archivo fue creado con "oasproj init" y será sobre-escrito con "oasproj update --force"`
	return
}

func AllTemplates(ctx *cli.Context, this_project project.Project) []error {
	errors := []error{}
	if ctx.Bool("with-daemon") {
		templateFileContent[".internal/root/usr/lib/systemd/system/"+this_project.Project_name+".service"] = systemd_service
	}
	for file_path, file_template := range templateFileContent {
		file_path = filepath.Join(ctx.Args().First(), file_path)
		dir := path.Dir(file_path)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				errors = append(errors, err)
			}
		}
		_, file_stat_err := os.Stat(file_path)
		file_exists := file_stat_err == nil
		create_only_file := createOnlyFiles[file_path]
		write_file := false
		if create_only_file {
			if !file_exists {
				write_file = true
			} else if ctx.Bool("force") {
				write_file = true
			}
		} else {
			write_file = true
		}
		if write_file {
			fmt.Println("Escribiendo " + file_path)
			if err := writeTemplate(file_path, file_template, this_project); err != nil {
				errors = append(errors, err)
			}
		}
	}
	return errors
}

func writeTemplate(file_path, file_template string, this_project project.Project) (err error) {
	var file *os.File
	if file, err = os.Create(file_path); err == nil {
		defer file.Close()
		var t *template.Template
		if t, err = template.New(file_path).Parse(file_template); err == nil {
			template_info := CreateDefaultTemplateInfo(this_project)
			err = t.Execute(file, &template_info)
		}
	}
	return
}
