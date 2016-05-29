package template

var scripts_after_install = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer después de instalar el paquete aquí:
#
echo After install
`
