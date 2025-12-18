package routes

import (
	"ainyx/internal/handler"
	"ainyx/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) {

	// AUTH ROUTES (public)
	auth := app.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	// Protected route
	auth.Get("/me", middleware.JWTProtected(), authHandler.Me)

	// USER ROUTES (protected)
	users := app.Group("/users", middleware.JWTProtected())
	users.Post("/", userHandler.CreateUser)
	users.Get("/:id", userHandler.GetUser)
	users.Get("/", userHandler.ListUsers)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}
