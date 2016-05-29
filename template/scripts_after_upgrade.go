package template

var scripts_after_upgrade = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer después de actualizar el paquete aquí:
#
echo After upgrade
`
