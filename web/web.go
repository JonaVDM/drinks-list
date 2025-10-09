package web

import (
	"embed"
	"io/fs"
)

//go:embed .output/**/*
var frontend embed.FS

var DistDirFS, _ = fs.Sub(frontend, ".output/public")
