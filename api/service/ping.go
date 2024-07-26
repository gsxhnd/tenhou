package service

import (
	"github.com/gsxhnd/tenhou/db"
	"github.com/gsxhnd/tenhou/utils"
)

type PingService interface {
	Ping() error
}

type pingService struct {
	logger utils.Logger
	db     *db.Database
}

func NewPingService(l utils.Logger, db *db.Database) PingService {
	return &pingService{
		logger: l,
		db:     db,
	}
}

func (p *pingService) Ping() error {
	if err := p.db.TenhouDB.Ping(); err != nil {
		p.logger.Errorf(err.Error())
		return err
	}

	p.logger.Infof("json db path: %s", p.db.TenhouJsonDB.Path())

	return nil
}
