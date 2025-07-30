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

type WrappedWriter struct {
	w        http.ResponseWriter
	b        []byte
	notFound *bool
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
	w.b = append(w.b, b...)
	return len(b), nil
}

func init() {
	sub, err := fs.Sub(buildDir, "build")
	if err != nil {
		panic(err)
	}
	ServeFS = HTMLFS{d: sub}

	Router.Use(middleware.StripSlashes)
	Router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		nf := false
		ww := WrappedWriter{w: w, notFound: &nf}
		http.FileServerFS(ServeFS).ServeHTTP(ww, r)
		b := ww.b
		if nf {
			www := WrappedWriter{w: w, notFound: nil}
			http.ServeFileFS(www, r, ServeFS, "index.html")
			b = www.b
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})
	Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, ServeFS, "index.html")
	})
}
