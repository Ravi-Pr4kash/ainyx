package main

import (
    "ainyx/config"
    "ainyx/internal/handler"
    "ainyx/internal/logger"
    "ainyx/internal/repository"
    "ainyx/internal/routes"
    "ainyx/internal/service"

    "github.com/gofiber/fiber/v2"
    "go.uber.org/zap"
)

func main() {
    // Init global logger
    logger.InitLogger()
    logger.Log.Info("Starting server...")

    // Load config (.env)
    cfg := config.LoadConfig()

    // Connect DB
    db, _, err := repository.ConnectDB(cfg)
    if err != nil {
        logger.Log.Fatal("Could not connect to DB", zap.Error(err))
    }

    // Initialize services
    userService := service.NewUserService(db)
    authService := service.NewAuthService(db)

    // Initialize handlers
    userHandler := handler.NewUserHandler(userService)
    authHandler := handler.NewAuthHandler(authService)

    // Create Fiber app
    app := fiber.New()

    // Register all routes
    routes.RegisterRoutes(app, userHandler, authHandler)

    // Start server
    logger.Log.Info("Server running on :3000")
    app.Listen(":3000")
}
