package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/gsxhnd/tenhou/db"
	"github.com/gsxhnd/tenhou/model"
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

		var startDate, endDate time.Time
		var fullData = make([]Paifu, 0)
		var dbData = make([]model.Log, 0)
		var notExistData = make([]Paifu, 0)

		for _, v := range htmlList {
			data, _ := ReadSingleFile(v)
			for _, d := range data {
				t, _ := time.Parse(time.RFC3339, d.Date)
				if startDate.IsZero() {
					startDate = t
				}
				if endDate.IsZero() {
					endDate = t
				}

				if startDate.After(t) {
					startDate = t
				}
				if endDate.Before(t) {
					endDate = t
				}
			}
			fullData = append(fullData, data...)
		}

		rows, err := db.TenhouDB.Query(
			"select id, log_id,game_type,game_date from tenhou where game_date >= ? and game_date <= ?",
			startDate.Format(time.RFC3339), endDate.Format(time.RFC3339),
		)
		if err != nil {
			fmt.Println(err)
		}

		for rows.Next() {
			var data = model.Log{}
			if err := rows.Scan(&data.Id, &data.LogId, &data.GameType, &data.GameDate); err != nil {
				logger.Errorw("GetLogInfo scan row error", "error", err)
				return err
			}
			dbData = append(dbData, data)
		}

		for _, v := range fullData {
			var exist bool = false
			for _, e := range dbData {
				if v.LogID == e.LogId {
					exist = true
				}
			}

			if !exist {
				notExistData = append(notExistData, v)
			}
		}

		if len(notExistData) == 0 {
			logger.Infof("not exist data in database is zero")
			return nil
		}

		logger.Infof("not exist data in database count is: %v", len(notExistData))
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
