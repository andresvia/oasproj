package template

var scripts_build = `# Variables de entorno disponibles durante construcción

# ${OAS_VERSION}      => contiene la version a construir
# ${OAS_PACKAGE_NAME} => nombre del paquete que proviene del archivo de descripción de proyecto

{{if eq (.Programming_language) "Go" -}}
go get -v ./...
go build -v ./...
{{- end}}
`
