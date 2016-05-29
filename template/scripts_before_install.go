package template

var scripts_before_install = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer antes de instalar el paquete aquí:
#
echo OAS before install
`
