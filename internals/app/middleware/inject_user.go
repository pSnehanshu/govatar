package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/internals/service/auth"
)

func (m *middlewares) InjectUser(c *fiber.Ctx) error {
	token := c.Cookies("access-token")

	user, err := auth.ValidateAccessToken(c.Context(), m.db, token)
	if err == nil {
		c.Locals("user", user)
	}

	return c.Next()
}
