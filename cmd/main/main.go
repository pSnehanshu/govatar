package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pSnehanshu/govatar/ent"
	"github.com/pSnehanshu/govatar/internals/app"
)

func main() {
	// Initializing .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	// Initialize database connection
	db, err := ent.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	// Determine the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "7180"
	}
	listenAddr := fmt.Sprintf(":%s", port)

	// Initialize and start the server (app)
	err = app.NewApp(db).Listen(listenAddr)
	log.Fatalf("Server ended with error: %v", err)
}
