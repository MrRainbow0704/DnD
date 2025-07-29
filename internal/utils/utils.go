package utils

import (
	"context"

	"github.com/MrRainbow0704/DnD/internal/config"
	database "github.com/MrRainbow0704/DnD/internal/db"
	"github.com/MrRainbow0704/DnD/internal/db/sqlc"
)

var (
	cnf = config.Get()
	db  = database.Get()
)

func SetUserAdmin(ctx context.Context, id int64) (sqlc.User, error) {
	return db.SetUserRole(ctx, RoleAdmin, id)
}

func RemoveUserAdmin(ctx context.Context, id int64) (sqlc.User, error) {
	return db.SetUserRole(ctx, RoleUser, id)
}
