package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/MrRainbow0704/DnD/api"
	"github.com/MrRainbow0704/DnD/internal/config"
	"github.com/MrRainbow0704/DnD/internal/log"
	"github.com/MrRainbow0704/DnD/internal/utils"
	"github.com/MrRainbow0704/DnD/internal/version"
	"github.com/MrRainbow0704/DnD/web"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
)

var (
	cnf = config.Get()
	v   = flag.Bool("v", false, "query version")
)

func main() {
	flag.Parse()
	if *v {
		fmt.Println(version.Get())
		return
	}

	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	r := chi.NewRouter()

	// Set essential middlewares
	chi_middleware.DefaultLogger = chi_middleware.RequestLogger(
		&chi_middleware.DefaultLogFormatter{
			Logger:  log.WebLogger,
			NoColor: true,
		},
	)
	r.Use(
		chi_middleware.SupressNotFound(r),
		chi_middleware.RequestID,
		chi_middleware.RealIP,
		chi_middleware.Logger,
		chi_middleware.Recoverer,
	)

	// Set additional middlewares
	r.Use(
		chi_middleware.CleanPath,
		chi_middleware.Timeout(60*time.Second),
	)

	// Mount API router
	r.Mount("/api/v1", api.Router)

	// Mount webserver router
	r.Mount("/", web.Router)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSON(w, http.StatusNotFound, nil, nil)
	})

	// Run the server
	s := &http.Server{
		Addr:    cnf.Address,
		Handler: r,
	}
	fmt.Printf("Running server on \"http://%s\"...", cnf.Address)
	return s.ListenAndServe()
}