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
type WrappedWriter struct {
	w        http.ResponseWriter
	notFound *bool
}

func (h HTMLFS) Open(name string) (fs.File, error) {
	if filepath.Ext(name) == "" {
		name += ".html"
	}
	return h.d.Open(name)
}

func (w WrappedWriter) WriteHeader(statusCode int) {
	if statusCode == http.StatusNotFound && w.notFound != nil {
		*w.notFound = true
	}
}

func (w WrappedWriter) Header() http.Header {
	return w.w.Header()
}

func (w WrappedWriter) Write(b []byte) (int, error) {
	if w.notFound != nil && *w.notFound {
		return 0, nil
	}
	return w.w.Write(b)
}

var (
	//go:embed all:build
	buildDir embed.FS
	Router   = chi.NewRouter()
	ServeFS  HTMLFS
)

func AnyHandler(w http.ResponseWriter, r *http.Request) {
	nf := false
	ww := WrappedWriter{w: w, notFound: &nf}
	http.FileServerFS(ServeFS).ServeHTTP(ww, r)
	if nf {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		http.ServeFileFS(w, r, ServeFS, "index.html")
	}
}

func init() {
	sub, err := fs.Sub(buildDir, "build")
	if err != nil {
		panic(err)
	}
	ServeFS = HTMLFS{d: sub}

	Router.Use(middleware.StripSlashes)
	Router.HandleFunc("/*", AnyHandler)
}
