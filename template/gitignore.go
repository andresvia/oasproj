package template

var gitignore = `# {{.Update}}
{{if eq (.Project.Programming_language) "Go" -}}

# Compiled Object files, Static and Dynamic libs (Shared Objects)
*.o
*.a
*.so

# Folders
_obj
_test

# Architecture specific extensions/prefixes
*.[568vq]
[568vq].out

*.cgo1.go
*.cgo2.c
_cgo_defun.c
_cgo_gotypes.go
_cgo_export.*

_testmain.go

*.exe
*.test
*.prof

{{.Project.Project_name}}
{{.Project.Project_name}}_*

{{- end}}

target
target-root
.vagrant
`
