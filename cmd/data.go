package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/tenhou_data.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
CREATE TABLE IF NOT EXISTS "tenhou"
(
    "id"        INTEGER NOT NULL UNIQUE,
    "log_id"    VARCHAR,
    "game_type" VARCHAR,
    "game_date" DATETIME,
    PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "tenhou_index_0"
    ON "tenhou" ("id", "log_id");
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}
	return db
}
