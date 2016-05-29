package template

var systemd_service = `# {{.ForceUpdate}}
[Unit]
Description=udpack
After=network.target

[Service]
ExecStart=/usr/bin/{{.Project.Project_name}}
Type=simple

[Install]
WantedBy=default.target
`
