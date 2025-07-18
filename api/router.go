package api

import (
	"net/http"

	"github.com/MrRainbow0704/DnD/api/handler"
	"github.com/MrRainbow0704/DnD/api/middleware"
	"github.com/MrRainbow0704/DnD/internal/utils"

	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

var Router = chi.NewRouter()

func init() {
	// Set middlewares
	Router.Use(chi_middleware.AllowContentType("application/json"))

	Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		allRoutes := []map[string]string{}
		chi.Walk(
			Router,
			func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
				allRoutes = append(
					allRoutes,
					map[string]string{
						"method": method,
						"route":  route,
					},
				)
				return nil
			},
		)
		utils.SendJSON(
			w,
			http.StatusOK,
			map[string]any{"routes": allRoutes},
			nil,
		)
	})
	Router.Get("/t", func(w http.ResponseWriter, r *http.Request) {
		utils.SendJSON(w, http.StatusOK, nil, nil)
	})
	Router.Post("/login", handler.Login)
	Router.Post("/user", handler.CreateUser)
	Router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middleware.Authenticator(utils.TokenAuth))

		r.Post("/logout", handler.Logout)
		r.Post("/character", handler.CreateCharacter)
		r.Get("/character/{id}", handler.GetCharacter)
	})
}
