package api

import (
	"fmt"
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
		)
	})
	Router.Post("/login", handler.Login)
	Router.Get("/users/{id}", handler.GetUser)
	Router.Post("/users", handler.CreateUser)
	Router.Get("/campaigns/{id}", handler.GetCampaign)
	Router.Get("/characters/{id}", handler.GetCharacter)
	Router.Group(func(r chi.Router) { // Authenticated routes
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(middleware.Authenticator(utils.TokenAuth))

		// r.Get("/me", handler.GetThisUser)
		// r.Patch("/me", handler.EditThisUser)
		// r.Delete("/me", handler.DeleteThisUser)
		r.Post("/logout", handler.Logout)
		r.Post("/campaigns", handler.CreateCampaign)
		// r.Patch("/campaigns/{id}", handler.EditCampaign)
		r.Post("/characters", handler.CreateCharacter)
		// r.Patch("/characters/{id}", handler.EditCharacter)
		Router.Group(func(ar chi.Router) { // Admin routes
			ar.Use(middleware.IsAdmin)

			// ar.Delete("/users/{id}", DeleteUser)
			// ar.Patch("/users/{id}", EditUser)
		})
		Router.Group(func(or chi.Router) { // Owner routes
			or.Use(middleware.IsOwner)

			// or.Delete("/characters/{id}", DeleteCharacter)
			// or.Patch("/characters/{id}", EditCharacter)
			// or.Delete("/campaigns/{id}", DeleteCampaign)
			// or.Patch("/campaigns/{id}", EditCampaign)
		})
	})
	Router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.ErrorJSON(
			w,
			http.StatusNotFound,
			map[string]error{"NOT_FOUND": fmt.Errorf("Pagina non trovata")},
		)
	})
}
