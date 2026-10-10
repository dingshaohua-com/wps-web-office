package static

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed files/*
var assets embed.FS

func Handler() http.Handler {
	files, err := fs.Sub(assets, "files")
	if err != nil {
		panic(err)
	}
	return http.FileServer(http.FS(files))
}
