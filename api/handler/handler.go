package handler

import (
	"net/http"

	"github.com/MrRainbow0704/DnD/internal/config"
	database "github.com/MrRainbow0704/DnD/internal/db"
	t "github.com/MrRainbow0704/DnD/internal/types"
	"github.com/MrRainbow0704/DnD/internal/utils"
	"github.com/go-chi/chi/v5"
)

type M = t.AnyMap
type E = t.ErrorMap

var (
	db  = database.Get() // Database instance
	cnf = config.Get()   // Config instance
)

// Keys for map access
const (
	msgKey       = "message"
	tokenKey     = "token"
	userKey      = "user"
	characterKey = "character"
	campaignKey  = "campaign"
)

// Error codes for API responses
const (
	jsonDecodeError    = "JSON_DECODE"
	emptyFieldError    = "EMPTY_FIELD"
	invalidParamError  = "INVALID_URL_PARAMETER"
	notFoundError      = "USER_NOT_FOUND"
	internalError      = "INTERNAL"
	credentialsError   = "INVALID_CREDENTIALS"
	alreadyExistsError = "USER_ALREADY_EXISTS"
)

func RootHandler(router chi.Router) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		routes := []M{}
		chi.Walk(
			router,
			func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
				routes = append(
					routes,
					M{"method": method, "route": route},
				)
				return nil
			},
		)
		utils.SendJSON(
			w,
			http.StatusOK,
			M{"routes": routes},
		)
	}
}
