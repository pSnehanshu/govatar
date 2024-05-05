package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pSnehanshu/govatar/ent"
)

func main() {

	// Initialize HTTP router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(attachDBCtx)

	// Define the logs
	r.Get("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(chi.URLParam(r, "id"))
		if err != nil {
			w.Write([]byte("Invalid UUID"))
			return
		}

		ctx := r.Context()
		db := ctx.Value(dbClientKey).(*ent.Client)

		user, err := db.User.Get(ctx, id)

		if err != nil {
			w.Write([]byte(fmt.Sprintf("User error: %v", err)))
			return
		}

		w.Write([]byte(fmt.Sprintf("Hello %s", user.Email)))
	})

	// Start server
	log.Println("Server listening")
	http.ListenAndServe(":3000", r)
}
