package utils

import (
	"github.com/MrRainbow0704/DnD/internal/config"
	"github.com/go-chi/jwtauth/v5"
)

var (
	cnf       = config.Get()
	TokenAuth *jwtauth.JWTAuth
)

func init() {
	TokenAuth = jwtauth.New("HS256", cnf.JWTKey, nil)
}
