package app

import (
	"fmt"

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

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		user, err := db.User.Get(c.Context(), id)

		if err != nil {
			return c.SendString(fmt.Sprintf("User error: %v", err))
		}

		return c.SendString(fmt.Sprintf("Hello %s %s", user.ID, user.CreatedAt))
	})
}
