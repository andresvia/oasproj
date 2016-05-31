package template

var vagrantfile = `# -*- mode: ruby -*-
# vi: set ft=ruby :

# {{.ForceUpdate}}

Vagrant.configure(2) do |config|
  config.vm.box = "centos/7"
  # config.vm.network "private_network", ip: "192.168.12.x"
  config.vm.provision "shell", inline: "rm -rf /tmp/target", run: "always"
  config.vm.provision "file", source: "target", destination: "/tmp/target", run: "always"
  config.vm.provision "shell", path: "scripts/vagrant-installer", run: "always"
  config.vm.provision "shell", inline: "rm -rf /tmp/target/", run: "always"
end
`
