package db

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/gsxhnd/tenhou/utils"
	_ "github.com/mattn/go-sqlite3"
	"go.etcd.io/bbolt"
)

type Database struct {
	TenhouDB     *sql.DB
	TenhouJsonDB *bbolt.DB
}

func NewDatabase(cfg *utils.Config, l utils.Logger) (*Database, error) {
	dbSource := cfg.TenhouDBPath
	db, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS "tenhou"
	(
		"id"        INTEGER NOT NULL UNIQUE,
		"log_id"    VARCHAR NOT NULL UNIQUE,
		"game_type" VARCHAR,
		"game_date" DATETIME,
		PRIMARY KEY ("id")
	);
	
	CREATE INDEX IF NOT EXISTS "tenhou_index_0"
		ON "tenhou" ("id", "log_id");
		`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}

	jsonDB, err := bbolt.Open(cfg.TenhouJsonDBPath, 0666, nil)
	if err != nil {
		return nil, err
	}

	return &Database{
		TenhouDB:     db,
		TenhouJsonDB: jsonDB,
	}, nil
}

var DBSet = wire.NewSet(NewDatabase)
