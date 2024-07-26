package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/utils"
)

type Middlewarer interface {
	RequestLog(ctx *fiber.Ctx) error
}
type middleware struct {
	logger utils.Logger
}

func NewMiddleware(l utils.Logger) Middlewarer {
	return &middleware{
		logger: l,
	}
}
