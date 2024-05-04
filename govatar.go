package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/pSnehanshu/govatar/ent"
)

func main() {
	client, err := ent.Open("postgres", "postgresql://postgres:password@127.0.0.1/govatar?sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// user, err := createUser(ctx, client)
	// if err != nil {
	// 	log.Fatalf("failed to create user: %v", err)
	// }

	// log.Printf("The created user is %v", user)
}

// func createUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("a8m").
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed creating user: %w", err)
// 	}
// 	log.Println("user was created: ", u)
// 	return u, nil
// }
