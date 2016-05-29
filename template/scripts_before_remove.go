package template

var scripts_before_remove = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer antes de remover el paquete aquí:
#
echo OAS before remove
`
