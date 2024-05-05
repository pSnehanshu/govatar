package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
)

func NewApp(db *ent.Client) *fiber.App {
	app := fiber.New()

	mountRoutes(app, db)

	return app
}
