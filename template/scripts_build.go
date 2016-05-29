package template

var scripts_build = `# {{.ForceUpdate}}
# Variables de entorno disponibles durante construcción

# ${OAS_VERSION}      => contiene la version a construir
# ${OAS_PACKAGE_NAME} => nombre del paquete que proviene del archivo de descripción de proyecto

{{if eq (.Project.Programming_language) "Go" -}}

# construir

go get -v ./...
go build -v ./...

# luego copiar el archivo compilado a su ubicación final

cp -v "{{.Project.Project_name}}" "target-root/usr/bin/{{.Project.Project_name}}"

{{- end}}
`
