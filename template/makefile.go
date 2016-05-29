package template

var makefile = `# {{.Update}}
all:
	bash .internal/build

test:
	bash .internal/test
`
