package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pSnehanshu/govatar/ent"
)

func mountRoutes(app *fiber.App, db *ent.Client) {
	// Define the logs
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.SendString("Invalid UUID")
		}

		user, err := db.User.Get(c.Context(), id)

		if err != nil {
			return c.SendString(fmt.Sprintf("User error: %v", err))
		}

		return c.SendString(fmt.Sprintf("Hello %s", user.Email))
	})
}
