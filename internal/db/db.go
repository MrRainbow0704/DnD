package db

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/MrRainbow0704/DnD/internal/config"
	"github.com/MrRainbow0704/DnD/internal/db/sqlc"
	"github.com/MrRainbow0704/DnD/internal/log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

var (
	cnf     = config.Get()
	queries *sqlc.Queries
)

func init() {
	ctx := context.Background()

	// Open connection to DB
	src := `file:` + cnf.DBPath + `?cache=shared&mode=rwc`
	conn, err := sql.Open("sqlite3", src)
	if err != nil {
		log.Panicf("Errore durante la connessione al database: %s", map[string]any{"source": src}, err)
		return
	}

	// Execute schema.sql
	if _, err := conn.ExecContext(ctx, ddl); err != nil {
		log.Panicf("Errore durante l'esecuzione dello schema: %s", nil, err)
	}

	queries = sqlc.New(conn)
}

func Get() *sqlc.Queries {
	return queries
}
