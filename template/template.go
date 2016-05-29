package template

import (
	"fmt"
	"github.com/andresvia/oasproj/project"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

var templateFileContent = map[string]string{
	"LICENSE":         license,
	"Makefile":        makefile,
	".internal/build": internal_build,
	"scripts/build":   scripts_build,
}

var createOnlyFiles = map[string]bool{
	"scripts/build": true,
}

func AllTemplates(base_path string, this_project project.Project, force bool) []error {
	errors := []error{}
	for file_path, file_template := range templateFileContent {
		file_path = filepath.Join(base_path, file_path)
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
			} else if force {
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
		if t, err = template.New("project").Parse(file_template); err == nil {
			err = t.Execute(file, &this_project)
		}
	}
	return
}
