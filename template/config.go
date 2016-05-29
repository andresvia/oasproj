package template

var templateFileContent = map[string]string{
	"LICENSE":                license,
	"Makefile":               makefile,
	".internal/build":        internal_build,
	"scripts/build":          scripts_build,
	"scripts/after-install":  scripts_after_install,
	"scripts/after-remove":   scripts_after_remove,
	"scripts/after-upgrade":  scripts_after_upgrade,
	"scripts/before-install": scripts_before_install,
	"scripts/before-remove":  scripts_before_remove,
	"scripts/before-upgrade": scripts_before_upgrade,
}

var createOnlyFiles = map[string]bool{
	"scripts/build":          true,
	"scripts/after-install":  true,
	"scripts/after-remove":   true,
	"scripts/after-upgrade":  true,
	"scripts/before-install": true,
	"scripts/before-remote":  true,
	"scripts/before-upgrade": true,
}
