package middleware

import "github.com/gofiber/fiber/v2"

func (m *middleware) RequestLog(ctx *fiber.Ctx) error {

	ctx.Next()
	var key = []interface{}{
		"method", ctx.Method(),
		"path", ctx.Path(),
		"client_id", ctx.IP(),
		"status", ctx.Response().StatusCode(),
	}
	m.logger.Infow("", key...)
	return nil

}
