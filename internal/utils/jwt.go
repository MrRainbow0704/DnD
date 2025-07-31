package utils

import (
	"github.com/go-chi/jwtauth/v5"
)

const (
	JWTRoleKey = "JWT.key.user_role"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

var TokenAuth = jwtauth.New("HS256", []byte(cnf.JWTKey), nil)
