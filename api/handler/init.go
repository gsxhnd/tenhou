package handler

import (
	"github.com/google/wire"
)

type Handler struct {
	PingHandler PingHandler
	LogHandler  LogHandler
}

var HandlerSet = wire.NewSet(
	NewPingHandler,
	NewLogHandler,
	wire.Struct(new(Handler), "*"),
)
