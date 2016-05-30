package template

var systemd_service = `# {{.ForceUpdate}}
[Unit]
Description={{.Project.Project_description}}
After=network.target

[Service]
ExecStart=/usr/bin/{{.Project.Project_name}}
Type=simple

[Install]
WantedBy=default.target
`
