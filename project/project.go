package project

import (
	"bytes"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

const metadata_file = ".oas.yml"

var showProjectTemplate = `Nombre del proyecto: {{.Project_name}}
Descripción del proyecto: {{.Project_description}}
Propósito del proyecto: {{.Project_purpose}}
Lenguaje de programación: {{.Programming_language}}
Marco de trabajo de proyecto: {{.Project_framework}}
Unidad organizacional: {{.Organizational_unit}}
Servicio: {{.Daemon}}
Dependencias del paquete: {{range $key, $value := .Package_dependencies}}
  {{$key}}{{range $value}}
  - {{.}}{{end}}{{end}}
Dependencias de compilación: {{range $key, $value := .Build_dependencies}}
  {{$key}}{{range $value}}
  - {{.}}{{end}}{{end}}`

type Project struct {
	Project_name         string              `yaml:project_name`
	Project_description  string              `yaml:project_description`
	Project_purpose      string              `yaml:project_purpose`
	Programming_language string              `yaml:programming_language`
	Project_framework    string              `yaml:project_framework`
	Organizational_unit  string              `yaml:organizational_unit`
	Package_dependencies map[string][]string `yaml:package_dependencies`
	Build_dependencies   map[string][]string `yaml:build_dependencies`
	Daemon               bool                `yaml:daemon`
}

func (p Project) String() (s string) {
	var err error
	var t *template.Template
	if t, err = template.New("project").Parse(showProjectTemplate); err == nil {
		var txt bytes.Buffer
		if err = t.Execute(&txt, &p); err == nil {
			s = txt.String()
		} else {
			s = "TEMPLATE_ERROR=Error ejecutando la plantilla: " + err.Error()
		}
	} else {
		s = "TEMPLATE_ERROR=Error en la plantilla: " + err.Error()
	}
	return
}

func LoadProject(path string) (p Project) {
	var err error
	var metadata_fd *os.File
	var metadata_content []byte
	metadata_path := getMetadataPath(path)
	if metadata_fd, err = os.Open(metadata_path); err == nil {
		defer metadata_fd.Close()
		if metadata_content, err = ioutil.ReadAll(metadata_fd); err == nil {
			err = yaml.Unmarshal(metadata_content, &p)
		}
	}
	if err != nil {
		panic("Error al leer archivo descriptor del proyecto " + metadata_path + ": " + err.Error())
	}
	return
}

func getMetadataPath(path string) (metadata_path string) {
	metadata_path = filepath.Join(path, metadata_file)
	return
}

func MetadataExists(path string) (exists bool) {
	if _, err := os.Stat(getMetadataPath(path)); os.IsNotExist(err) {
		exists = false
	} else {
		exists = true
	}
	return
}

func New(ctx *cli.Context) (p Project) {
	p.Project_name = ctx.String("name")
	p.Project_description = ctx.String("desc")
	p.Project_purpose = ctx.String("purpose")
	p.Programming_language = ctx.String("language")
	p.Project_framework = ctx.String("framework")
	p.Organizational_unit = ctx.String("orgunit")
	return
}

func (p Project) WriteFile(base string) (err error) {
	file_path := getMetadataPath(base)
	var file *os.File
	var buf []byte
	if buf, err = yaml.Marshal(p); err == nil {
		if file, err = os.Create(file_path); err == nil {
			defer file.Close()
			_, err = file.Write(buf)
		}
	}
	return
}
