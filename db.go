package main

import (
	"log"

	"github.com/pSnehanshu/govatar/ent"
)

// Get instance of Ent Client for database operation. It will exit the program if couldn't connect.
func getDBClient() *ent.Client {
	db, err := ent.Open("postgres", "postgresql://postgres:password@127.0.0.1/govatar?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return db
}
