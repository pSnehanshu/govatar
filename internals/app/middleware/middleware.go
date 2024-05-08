package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
)

type middlewares struct {
	db *ent.Client
}

func New(db *ent.Client) middlewares {
	return middlewares{db}
}

type MiddlewareFunc func(*fiber.Ctx) error
