package db

import (
	"github.com/google/wire"
)

type Database struct {
}

func NewDatabase() (*Database, error) {
	// func NewDatabase(cfg *utils.Config, l utils.Logger) (*Database, error) {
	return &Database{}, nil
}

var DBSet = wire.NewSet(NewDatabase)
