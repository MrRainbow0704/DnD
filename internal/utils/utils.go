package utils

import (
	"context"

	"github.com/MrRainbow0704/DnD/internal/config"
	database "github.com/MrRainbow0704/DnD/internal/db"
	"github.com/MrRainbow0704/DnD/internal/db/sqlc"
	t "github.com/MrRainbow0704/DnD/internal/types"
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

func PrepareCharacter(ctx context.Context, c sqlc.Character) (t.Character, error) {
	clsLevel := make([]t.ClassLevel, 0)
	u, err := db.GetUser(ctx, c.Owner)
	if err != nil {
		return t.Character{}, err
	}

	n := t.Character{}
	n.Owner = c.Owner
	n.CharName = c.Name
	n.OtherProfs = c.Proficencies.(string)
	n.Misc = t.Misc{
		ClassLevel: clsLevel,
		Background: c.Background.(string),
		PlayerName: u.Name,
	}
	return n, nil
}
