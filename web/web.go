package web

import (
	"embed"
	"io/fs"
)

//go:embed dist
var frontend embed.FS

var DistDirFS, _ = fs.Sub(frontend, "dist")
