package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func main() {
	db := getDBClient()
	defer db.Close()

	app := fiber.New()

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

	// Start server
	app.Listen(":3000")
}
