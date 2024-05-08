package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
)

func (m *middlewares) AuthGuard(c *fiber.Ctx) error {
	if _, ok := c.Locals("user").(*ent.User); ok {
		return c.Next()
	}

	return c.SendStatus(401)
}
