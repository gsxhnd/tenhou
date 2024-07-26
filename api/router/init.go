package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gsxhnd/tenhou/api/handler"
	"github.com/gsxhnd/tenhou/api/middleware"
)

type Router interface {
	Run() error
}

type router struct {
	// cfg *utils.Config
	app *fiber.App
	h   handler.Handler
	m   middleware.Middlewarer
}

func NewRouter(m middleware.Middlewarer, h handler.Handler) (Router, error) {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes:     true,
		DisableStartupMessage: false,
		Prefork:               false,
	})

	return &router{
		app: app,
		h:   h,
		m:   m,
	}, nil
}

func (r *router) Run() error {
	api := r.app.Group("/api")
	api.Get("/ping", r.h.PingHandler.Ping)

	r.app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c.Request().URI())
		return c.SendStatus(404) // => 404 "Not Found"
	})

	return r.app.Listen(":8080")
}
