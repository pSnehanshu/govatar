package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pSnehanshu/govatar/ent"
	"github.com/pSnehanshu/govatar/ent/user"
)

// valid signing method
var valid_sigining_methods = []string{"HS256"}
var login_jwt_secret = []byte(os.Getenv("LOGIN_JWT_SECRET_KEY"))

func GenerateAccessToken(user *ent.User) (string, time.Time, error) {
	// Generate JWT token for session
	expires := time.Now().Add(3 * 24 * time.Hour)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.RegisteredClaims{
		Issuer:    "govatar",
		ExpiresAt: jwt.NewNumericDate(expires),
		Subject:   user.ID,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}).SignedString(login_jwt_secret)

	return token, expires, err
}

func ValidateAccessToken(ctx context.Context, db *ent.Client, token string) (*ent.User, error) {
	if token != "" {
		// Validate token
		parsed_token, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if !contains(valid_sigining_methods, t.Method.Alg()) {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return login_jwt_secret, nil
		}, jwt.WithValidMethods(valid_sigining_methods))

		if err == nil && parsed_token.Valid {
			// Fetch the user
			if user_id, err := parsed_token.Claims.GetSubject(); err == nil {
				query := db.User.
					Query().
					Select(user.FieldID, user.FieldName, user.FieldCreatedAt).
					Where(user.ID(user_id))

				if user, err := query.Only(ctx); err == nil {
					return user, nil
				} else {
					return nil, err
				}
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no token submitted")
	}
}

func contains(slice []string, element string) bool {
	for _, item := range slice {
		if item == element {
			return true
		}
	}
	return false
}
