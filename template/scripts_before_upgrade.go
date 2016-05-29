package template

var scripts_before_upgrade = `#!/bin/bash
# {{.ForceUpdate}}
set -eu
# cosas para hacer antes de actulizar el paquete aquí:
#
echo Before upgrade
`
