package db

import (
	"testing"

	"github.com/gsxhnd/tenhou/utils"
	"github.com/stretchr/testify/assert"
)

var db *Database

func init() {
	var cfg, _ = utils.NewConfig()
	cfg.TenhouDBPath = "../data/tenhou_data.db"
	cfg.TenhouJsonDBPath = "../data/tenhou_json.db"
	var logger = utils.NewLogger(cfg)
	logger.Infof("init logger success")

	initDB, err := NewDatabase(cfg, logger)
	if err != nil {
		panic(err)
	}

	err = initDB.TenhouDB.Ping()
	if err != nil {
		panic(err)
	}

	db = initDB
}

func TestDatabase_GetLogInfoByLogId(t *testing.T) {
	tests := []struct {
		name  string
		logID string
	}{
		{"test", "2009022323gm-00b9-0000-881cb83e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db.GetLogInfoByLogId(tt.logID)
		})
	}
}

func TestDatabase_GetLogInfoList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := db.GetLogInfoList()
			assert.Nil(t, err)
			assert.Equal(t, 20, len(got))
		})
	}
}
