package project

func (p *Project) GetSystemdUnitOrder() (s string) {
	if p.Systemd_unit_after != "" {
		s = "After=" + p.Systemd_unit_after
		s += "\nRequires=" + p.Systemd_unit_after
	} else {
		s = "After=network.target"
	}
	return
}
