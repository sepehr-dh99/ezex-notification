// Package embed defines global embed paths
package embed

import "embed"

// TemplateFS defines email markdown embed
//
//go:embed assets/template/email/*.md
var TemplateFS embed.FS
