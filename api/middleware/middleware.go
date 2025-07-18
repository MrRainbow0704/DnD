
// Package middleware is a subpackage of the api package which provides custom
// middleware for the parent api package.
package middleware

import (
	"context"
	"net/http"

	"github.com/MrRainbow0704/DnD/internal/utils"
	"github.com/go-chi/jwtauth/v5"
)

func Authenticator(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())

			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			if token == nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			claims := token.PrivateClaims()
			r = r.WithContext(
				context.WithValue(
					context.WithValue(
						r.Context(),
						"USER_ID",
						claims[utils.JWTIDKey],
					),
					"USER_ROLE",
					claims[utils.JWTRoleKey],
				),
			)
			next.ServeHTTP(w, r)
		})
	}
}
