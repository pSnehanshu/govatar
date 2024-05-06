package app

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/pSnehanshu/govatar/ent"
)

//go:embed views
var viewsfs embed.FS

func NewApp(db *ent.Client) *fiber.App {
	engine := html.NewFileSystem(http.FS(viewsfs), ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layouts/root",
	})

	mountRoutes(app, db)

	return app
}
