package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/api/service"
)

type PingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type pingHandle struct {
	validator *validator.Validate
	svc       service.PingService
}

func NewPingHandler(svc service.PingService) PingHandler {
	return &pingHandle{
		svc: svc,
	}
}

// @Description  ping serivce working, db connect
// @Produce      json
// @Success      200
// @Router       /ping [get]
func (h *pingHandle) Ping(ctx *fiber.Ctx) error {
	err := h.svc.Ping()
	if err != nil {
		return ctx.Status(400).SendString(err.Error())
	}
	return ctx.Status(200).SendString("pong")
}
