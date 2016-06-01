package template

var vagrantfile = `# -*- mode: ruby -*-
# vi: set ft=ruby :

# {{.ForceUpdate}}

Vagrant.configure(2) do |config|
  config.vm.box = "centos/7"
  # config.vm.hostname = "name.192.168.12.x.xip.io"
  # config.vm.network "private_network", ip: "192.168.12.x"
  # config.vm.network "forwarded_port", guest: 80, host: 8080
  config.vm.provision "shell", inline: "rm -rf /tmp/target", run: "always"
  config.vm.provision "file", source: "target", destination: "/tmp/target", run: "always"
  config.vm.provision "shell", path: "scripts/vagrant-installer", run: "always"
  config.vm.provision "shell", inline: "rm -rf /tmp/target/", run: "always"
end
`
