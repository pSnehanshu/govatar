package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
)

func mountRoutes(app *fiber.App, db *ent.Client) {
	// app.Get("/avatar/:hash", func(c *fiber.Ctx) error {
	// 	//
	// })

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("views/signup", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("views/login", fiber.Map{})
	})
}
