package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pSnehanshu/govatar/internals/app"
	"github.com/pSnehanshu/govatar/internals/database"
)

func main() {
	db := database.GetClient()
	defer db.Close()

	app := app.NewApp(db)

	log.Fatalf("Server ended with error: %v", app.Listen(":3000"))
}
