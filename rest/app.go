package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	listenAddr = "127.0.0.1:8080"
	baseURL    = "/api"
)

var app *fiber.App

func NewApp() error {
	var appCfg = fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
	}

	app = fiber.New(appCfg)
	app.Use(logger.New())
	initRoutes()

	if err := app.Listen(listenAddr); err != nil {
		return err
	}

	return nil
}

func Shutdown() error {
	return app.Shutdown()
}

func ping(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("pong")
}

func initRoutes() {
	base := app.Group(baseURL)
	base.Get("/ping", ping)
}
