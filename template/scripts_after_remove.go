package template

var scripts_after_remove = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer después de remover el paquete aquí:
#
echo After remove
`
