package db

import (
	"github.com/gsxhnd/tenhou/model"
)

const SelectLogInfo = `select * from tenhou where log_id = ?;`
const SelectLogList = `select * from tenhou limit 20;`

func (db *Database) GetLogInfoByLogId(log string) (*model.Log, error) {
	var data = model.Log{}
	rows, err := db.TenhouDB.Query(SelectLogInfo, log)
	if err != nil {
		db.logger.Errorw("GetLogInfo select error", "error", err.Error())
		return nil, err
	}

	if rows.Next() {
		if err := rows.Scan(&data.Id, &data.LogId, &data.GameType, &data.GameDate); err != nil {
			db.logger.Errorw("GetLogInfo scan row error", "error", err)
			return nil, err
		}
	}
	return &data, nil
}

func (db *Database) GetLogInfoList() ([]model.Log, error) {
	var list = make([]model.Log, 0)
	rows, err := db.TenhouDB.Query(SelectLogList)
	if err != nil {
		db.logger.Errorw("GetLogInfo select error", "error", err.Error())
		return nil, err
	}

	for rows.Next() {
		var data = model.Log{}
		if err := rows.Scan(&data.Id, &data.LogId, &data.GameType, &data.GameDate); err != nil {
			db.logger.Errorw("GetLogInfo scan row error", "error", err)
			return nil, err
		}
		list = append(list, data)
	}

	return list, nil
}
