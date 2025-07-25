package handler

import (
	"github.com/MrRainbow0704/DnD/internal/config"
	database "github.com/MrRainbow0704/DnD/internal/db"
)

type M map[string]any   // Alias for map[string]any
type E map[string]error // Alias for map[string]error

var (
	db  = database.Get() // Database instance
	cnf = config.Get()   // Config instance
)

const (
	msgKey       = "message"
	tokenKey     = "token"
	userKey      = "user"
	characterKey = "character"
	campaignKey  = "campaign"
)

const (
	jsonDecodeError    = "JSON_DECODE"
	emptyFieldError    = "EMPTY_FIELD"
	invalidParamError  = "INVALID_URL_PARAMETER"
	notFoundError      = "USER_NOT_FOUND"
	internalError      = "INTERNAL"
	credentialsError   = "INVALID_CREDENTIALS"
	alreadyExistsError = "USER_ALREADY_EXISTS"
)
