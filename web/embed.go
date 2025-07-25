package web

import (
	"embed"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTMLFS struct {
	d fs.FS
}

func (h HTMLFS) Open(name string) (fs.File, error) {
	if filepath.Ext(name) == "" {
		name += ".html"
	}
	return h.d.Open(name)
}

//go:embed all:build
var buildDir embed.FS
var (
	Router  = chi.NewRouter()
	ServeFS HTMLFS
)

func init() {
	sub, err := fs.Sub(buildDir, "build")
	if err != nil {
		panic(err)
	}
	ServeFS = HTMLFS{d: sub}

	Router.Use(middleware.StripSlashes)
	Router.Handle("/*", http.FileServerFS(ServeFS))
}
