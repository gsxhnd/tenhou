package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/api/handler"
	"github.com/gsxhnd/tenhou/api/middleware"
	"github.com/gsxhnd/tenhou/utils"
)

type Router interface {
	Run() error
}

type router struct {
	cfg *utils.Config
	app *fiber.App
	h   handler.Handler
	m   middleware.Middlewarer
}

// @title           Tenhou API
// @version         1
// @description     This is a sample server celler server.
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host      localhost:8080
// @BasePath  /api/v1
// @securityDefinitions.basic  BasicAuth
// @externalDocs.description  OpenAPI
func NewRouter(cfg *utils.Config, m middleware.Middlewarer, h handler.Handler) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     cfg.Mode == "dev",
		DisableStartupMessage: cfg.Mode == "prod",
		Prefork:               false,
	})

	return &router{
		cfg: cfg,
		app: app,
		h:   h,
		m:   m,
	}, nil
}

func (r *router) Run() error {
	r.app.Use(r.m.RequestLog)

	api := r.app.Group("/api/v1")
	api.Get("/ping", r.h.PingHandler.Ping)
	// api.Get("/paifu")
	// api.Get("/paifu/:log_id")

	r.app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Request().URI())
		return c.SendStatus(404)
	})

	return r.app.Listen(r.cfg.Listen)
}
