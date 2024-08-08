package main

import (
	"github.com/gsxhnd/tenhou/db"
	"github.com/gsxhnd/tenhou/utils"
	"github.com/urfave/cli/v2"
	"go.etcd.io/bbolt"
)

var storeJsonCmd = &cli.Command{
	Name:  "store_json",
	Usage: "",
	Action: func(ctx *cli.Context) error {
		cfg, _ := utils.NewConfig()
		logger := utils.NewLogger(cfg)
		db, _ := db.NewDatabase(cfg, logger)
		db.TenhouJsonDB.Batch(func(tx *bbolt.Tx) error {
			// tx.Bucket([]byte{}).
			return nil
		})
		return nil
	},
}
