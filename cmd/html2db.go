package main

import (
	"io/fs"
	"path/filepath"

	"github.com/gsxhnd/tenhou/db"
	"github.com/gsxhnd/tenhou/utils"
	"github.com/urfave/cli/v2"
)

var fullHtml2DB = &cli.Command{
	Name:  "full_html_to_db",
	Usage: "convert tenhou day data to csv",
	Action: func(ctx *cli.Context) error {
		var inputPath = "./data/tenhou_zip"
		var htmlList = make([]string, 0)
		cfg, _ := utils.NewConfig()
		logger := utils.NewLogger(cfg)
		db, _ := db.NewDatabase(cfg, logger)

		filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
			htmlList = append(htmlList, path)
			return nil
		})

		for _, v := range htmlList {
			data, _ := ReadSingleFile(v)
			tx, _ := db.TenhouDB.Begin()
			stmt, _ := tx.Prepare("insert into tenhou(log_id, game_type, game_date) values(?, ?,?)")
			defer stmt.Close()
			for _, d := range data {
				if _, err := stmt.Exec(d.LogID, d.GameType, d.Date); err != nil {
					logger.Errorw("stmt insert error", "error", err.Error())
				}
			}
			tx.Commit()
		}

		return nil
	},
}

var recentHtml2DB = &cli.Command{
	Name:  "recent_html_to_db",
	Usage: "convert tenhou day data to csv",
	Action: func(ctx *cli.Context) error {
		var inputPath = "./data/tencent_html"
		var htmlList = make([]string, 0)
		// config, _ := utils.NewConfig()
		// db, _ := db.NewDatabase(config, nil)

		filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
			htmlList = append(htmlList, path)
			return nil
		})

		var fullData = make([]Paifu, 0)
		for _, v := range htmlList {
			data, _ := ReadSingleFile(v)
			for _, d := range data {
				fullData = append(fullData, d)
			}
		}

		return nil
	},
}
