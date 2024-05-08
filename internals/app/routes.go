package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
)

type PostLoginBody struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func mountRoutes(app *fiber.App, db *ent.Client) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("views/signup", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("views/login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var body PostLoginBody

		if errs := validateReq(&body, c); len(errs) > 0 {
			return sendError(fiber.StatusBadRequest, c, &errs)
		}

		log.Printf("Body: %s %s", body.Email, body.Password)
		return c.SendStatus(fiber.StatusOK)
	})
}
