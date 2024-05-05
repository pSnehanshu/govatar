package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/pSnehanshu/govatar/ent"
	"github.com/pSnehanshu/govatar/internals/app"
)

func main() {
	// Initialize database connection
	db, err := ent.Open("postgres", "postgresql://postgres:password@127.0.0.1/govatar?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	// Initialize and start the server (app)
	err = app.NewApp(db).Listen(":3000")
	log.Fatalf("Server ended with error: %v", err)
}
