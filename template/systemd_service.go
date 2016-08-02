package template

var systemd_service = `# {{.ForceUpdate}}
[Unit]
Description={{.Project.Project_description}}
{{.Project.GetSystemdUnitOrder}}

[Service]
ExecStart=/usr/bin/{{.Project.Project_name}}
Type=simple

[Install]
WantedBy=default.target
`
