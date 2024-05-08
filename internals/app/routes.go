package app

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pSnehanshu/govatar/ent"
	"github.com/pSnehanshu/govatar/ent/email"
	"github.com/pSnehanshu/govatar/internals/app/middleware"
	"github.com/pSnehanshu/govatar/internals/service/auth"
	"golang.org/x/crypto/bcrypt"
)

type PostLoginBody struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func mountRoutes(app *fiber.App, db *ent.Client) {
	appMiddlewares := middleware.New(db)

	app.Use(appMiddlewares.InjectUser)

	app.Get("/", func(c *fiber.Ctx) error {
		log.Printf("User: %v\n", c.Locals("user"))

		return c.Render("views/index", fiber.Map{})
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.Render("views/signup", fiber.Map{})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("views/login", fiber.Map{})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var body PostLoginBody

		if errs := validateReq(&body, c); len(errs) > 0 {
			return sendError(fiber.StatusBadRequest, c, &errs)
		}

		// Find email
		email, err := db.Email.Query().Where(email.Email(body.Email)).Only(c.Context())
		if err != nil {
			log.Println("email not found")
			return sendError(fiber.StatusUnauthorized, c, &[]FieldError{{
				Error:   true,
				Message: "Failed to login",
			}})
		}

		if !email.IsVerified {
			return sendError(fiber.StatusUnauthorized, c, &[]FieldError{{
				Error:       true,
				Message:     "Email not verified",
				FailedField: "email",
				Value:       body.Email,
			}})
		}

		// Find associated user
		user, err := db.Email.QueryUser(email).Only(c.Context())
		if err != nil {
			log.Println("user not found for email", email.Email)
			return sendError(fiber.StatusUnauthorized, c, &[]FieldError{{
				Error:   true,
				Message: "Failed to login",
			}})
		}

		// Check password
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
		if err != nil {
			log.Println("incorrect password")
			return sendError(fiber.StatusUnauthorized, c, &[]FieldError{{
				Error:   true,
				Message: "Failed to login",
			}})
		}

		// Generate JWT token for session
		token, expires, err := auth.GenerateAccessToken(user)

		if err != nil {
			log.Println(err)
			return sendError(fiber.StatusInternalServerError, c, &[]FieldError{{
				Error:   true,
				Message: "Something went wrong!",
			}})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "access-token",
			Value:    token,
			Expires:  expires,
			HTTPOnly: true,
			SameSite: "Strict",
		})

		return c.SendStatus(fiber.StatusOK)
	})

	app.Get("/logout", appMiddlewares.AuthGuard, func(c *fiber.Ctx) error {
		c.ClearCookie("access-token")
		return c.Redirect("/")
	})
}
