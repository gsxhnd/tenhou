package service

type PingService interface{}

type pingService struct {
	// logger utils.Logger
	// db     *db.Database
}

func NewPingService() PingService {
	return &pingService{
		// logger: l,
		// db:     db,
	}
}
