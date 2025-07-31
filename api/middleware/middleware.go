// Package middleware is a subpackage of the api package which provides custom
// middleware for the parent api package.
package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	database "github.com/MrRainbow0704/DnD/internal/db"
	t "github.com/MrRainbow0704/DnD/internal/types"
	"github.com/MrRainbow0704/DnD/internal/utils"
	"github.com/go-chi/jwtauth/v5"
)

type E = t.ErrorMap

const (
	CtxUserID   t.CtxKey = "context.key.user_id"
	CtxUserRole t.CtxKey = "context.key.user_role"
)

const (
	internalError   = "INTERNAL"
	emptyFieldError = "EMPTY_FIELD"
	authError       = "UNAUTHORIZED"
)

var db = database.Get()

func Authenticator(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _, err := jwtauth.FromContext(r.Context())
			if err != nil {
				utils.ErrorJSON(
					w,
					http.StatusInternalServerError,
					E{internalError: fmt.Errorf("internal server error")},
				)
				return
			}
			if token == nil {
				utils.ErrorJSON(
					w,
					http.StatusUnauthorized,
					E{emptyFieldError: fmt.Errorf("empty token")},
				)
				return
			}
			claims := token.PrivateClaims()
			ctx := context.WithValue(r.Context(), CtxUserID, token.Subject())
			ctx = context.WithValue(ctx, CtxUserRole, claims[utils.JWTRoleKey])
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(CtxUserRole).(string)
		if !ok {
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("unauthorized")},
			)
			return
		}

		if role != utils.RoleAdmin {
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("role %s does not have permission to access this resource", role)},
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func IsOwner(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := r.Context().Value(CtxUserID).(int64)
		if !ok {
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("unauthorized")},
			)
			return
		}

		itemID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("unauthorized")},
			)
			return
		}

		var ownerID int64
		switch strings.Split(r.URL.Path, "/")[0] {
		case "campaigns":
			c, err := db.GetCampaign(r.Context(), int64(itemID))
			if err != nil {
				utils.ErrorJSON(
					w,
					http.StatusUnauthorized,
					E{authError: fmt.Errorf("unauthorized")},
				)
				return
			}
			ownerID = c.Master
		case "characters":
			c, err := db.GetCharacter(r.Context(), int64(itemID))
			if err != nil {
				utils.ErrorJSON(
					w,
					http.StatusUnauthorized,
					E{authError: fmt.Errorf("unauthorized")},
				)
				return
			}
			ownerID = c.Owner
		default:
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("unauthorized")},
			)
			return
		}

		if id != ownerID {
			utils.ErrorJSON(
				w,
				http.StatusUnauthorized,
				E{authError: fmt.Errorf("you are not the owner of this resource")},
			)
			return
		}

		next.ServeHTTP(w, r)
	})
}
