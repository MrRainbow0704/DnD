package api

import (
	"fmt"
	"net/http"

	h "github.com/MrRainbow0704/DnD/api/handler"
	m "github.com/MrRainbow0704/DnD/api/middleware"
	"github.com/MrRainbow0704/DnD/internal/utils"

	"github.com/go-chi/chi/v5"
	cm "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
)

var Router = chi.NewRouter()

func init() {
	// Set middlewares
	Router.Use(cm.AllowContentType("application/json"))

	Router.Get("/", h.RootHandler(Router))
	Router.Post("/login", h.Login)
	Router.With(m.ExtractID).Get("/users/{id}", h.GetUser)
	Router.Post("/users", h.CreateUser)
	Router.With(m.ExtractID).Get("/campaigns/{id}", h.GetCampaign)
	Router.With(m.ExtractID).Get("/characters/{id}", h.GetCharacter)
	Router.Group(func(r chi.Router) { // Authenticated routes
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(m.Authenticator(utils.TokenAuth))

		r.Get("/me", h.GetThisUser)
		r.Patch("/me", h.EditThisUser)
		r.Delete("/me", h.DeleteThisUser)
		r.Post("/logout", h.Logout)
		r.Post("/campaigns", h.CreateCampaign)
		// r.With(m.ExtractID).Patch("/campaigns/{id}", h.EditCampaign)
		r.Post("/characters", h.CreateCharacter)
		// r.With(m.ExtractID).Patch("/characters/{id}", h.EditCharacter)
		Router.Group(func(ar chi.Router) { // Admin routes
			ar.Use(m.IsAdmin)

			// ar.With(m.ExtractID).Delete("/users/{id}", h.DeleteUser)
			// ar.With(m.ExtractID).Patch("/users/{id}", h.EditUser)
		})
		Router.Group(func(or chi.Router) { // Owner routes
			or.Use(m.IsOwner)

			// or.With(m.ExtractID).Delete("/characters/{id}", h.DeleteCharacter)
			// or.With(m.ExtractID).Patch("/characters/{id}", h.EditCharacter)
			// or.With(m.ExtractID).Delete("/campaigns/{id}", h.DeleteCampaign)
			// or.With(m.ExtractID).Patch("/campaigns/{id}", h.EditCampaign)
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
