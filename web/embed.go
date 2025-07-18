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

var (
	//go:embed all:build
	buildDir embed.FS
	Router   = chi.NewRouter()
)

func init() {
	sub, err := fs.Sub(buildDir, "build")
	if err != nil {
		panic(err)
	}
	Router.Use(middleware.StripSlashes)
	Router.Handle("/*", http.FileServerFS(HTMLFS{d: sub}))
}