package handler

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/api/service"
)

type PingHandler interface {
	Ping(ctx *fiber.Ctx) error
}

type pingHandle struct {
	// logger    utils.Logger
	validator *validator.Validate
	svc       service.PingService
}

func NewPingHandler(svc service.PingService) PingHandler {
	return &pingHandle{
		// logger: l,
		svc: svc,
	}
}

func (h *pingHandle) Ping(ctx *fiber.Ctx) error {
	fmt.Println("ping")
	return ctx.Status(200).SendString("pong")
}
