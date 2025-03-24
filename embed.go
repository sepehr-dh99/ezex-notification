package embed

import "embed"

//go:embed assets/template/email/*.md
var TemplateFS embed.FS
