package template

var scripts_vagrant_installer = `# {{.ForceUpdate}}
# este instalador es válido sólo para ambientes Vagrant o de pruebas locales
sudo yum --nogpgcheck install -y /tmp/target/*.rpm
`
