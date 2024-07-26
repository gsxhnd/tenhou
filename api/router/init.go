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

// @title           Swagger Example API
// @version         2
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// ShowAccount godoc
// @Summary      Show an account


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
	api := r.app.Group("/api")

	api.Get("/ping", r.h.PingHandler.Ping)
	api.Get("/paifu")
	api.Get("/paifu/:log_id")

	r.app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Request().URI())
		return c.SendStatus(404)
	})

	return r.app.Listen(r.cfg.Listen)
}
