package main

import (
	"fmt"
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
		var inputPath = "./data/tenhou_html"
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
		var inputPath = "./data/tenhou_html"
		var htmlList = make([]string, 0)
		cfg, _ := utils.NewConfig()
		logger := utils.NewLogger(cfg)
		db, _ := db.NewDatabase(cfg, logger)

		filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
			if filepath.Ext(path) == ".html" {
				htmlList = append(htmlList, path)
			}
			return nil
		})

		var fullData = make([]Paifu, 0)
		var notExistData = make([]Paifu, 0)
		for _, v := range htmlList {
			data, _ := ReadSingleFile(v)
			fullData = append(fullData, data...)
		}

		for _, p := range fullData {
			var exist bool
			rows, err := db.TenhouDB.Query("SELECT EXISTS (SELECT 1 FROM tenhou WHERE log_id = ?) as exist", p.LogID)
			if err != nil {
				fmt.Println(err)
			}
			if rows.Next() {
				rows.Scan(&exist)
			}
			if !exist {
				notExistData = append(notExistData, p)
			}
		}

		fmt.Println(len(notExistData))

		tx, _ := db.TenhouDB.Begin()
		stmt, _ := tx.Prepare("insert into tenhou(log_id, game_type, game_date) values(?, ?,?)")
		defer stmt.Close()
		for _, d := range notExistData {
			if _, err := stmt.Exec(d.LogID, d.GameType, d.Date); err != nil {
				logger.Errorw("stmt insert error", "error", err.Error())

			}
		}
		if err := tx.Commit(); err != nil {
			tx.Rollback()
		}

		return nil
	},
}
