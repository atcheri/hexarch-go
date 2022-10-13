// Package docs provides utilities for api documentation generation
package docs

import (
	"embed"
)

//nolint:revive
//go:embed **
var Swagger embed.FS
