package main

import (
	"context"
	"log"
	"net/http"

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

type key uint8

const (
	dbClientKey key = 5
)

func attachDBCtx(next http.Handler) http.Handler {
	db := getDBClient()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), dbClientKey, db)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
