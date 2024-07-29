package service

import (
	"github.com/gsxhnd/tenhou/db"
	"github.com/gsxhnd/tenhou/model"
	"github.com/gsxhnd/tenhou/utils"
)

type LogService interface {
	GetLogInfoByLogId(logId string) (*model.Log, error)
	GetLogInfoList(page model.Pagination) ([]model.Log, error)
}

type logService struct {
	logger utils.Logger
	db     *db.Database
}

func NewLogService(l utils.Logger, db *db.Database) LogService {
	return &logService{
		logger: l,
		db:     db,
	}
}

func (p *logService) GetLogInfoByLogId(logId string) (*model.Log, error) {
	logInfo, err := p.db.GetLogInfoByLogId(logId)
	return logInfo, err
}

func (p *logService) GetLogInfoList(page model.Pagination) ([]model.Log, error) {
	logInfo, err := p.db.GetLogInfoList()
	return logInfo, err
}
