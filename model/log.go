package model

import "time"

// Log Data info
// @Description Log Data info
type Log struct {
	Id       uint      `json:"id"`        //id
	LogId    string    `json:"log_id"`    // log id
	GameType string    `json:"game_type"` // game type
	GameDate time.Time `json:"game_date"` // game start date
}
