package model

import "time"

type Log struct {
	Id       uint      `json:"id"`
	LogId    string    `json:"log_id"`
	GameType string    `json:"game_type"`
	GameDate time.Time `json:"game_date"`
}
