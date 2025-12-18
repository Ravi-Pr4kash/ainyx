package middleware

import (
    "os"

    "github.com/gofiber/fiber/v2"
    jwtware "github.com/gofiber/contrib/jwt"
)

func JWTProtected() fiber.Handler {
    secret := os.Getenv("JWT_SECRET")

    return jwtware.New(jwtware.Config{
        SigningKey: jwtware.SigningKey{
            Key: []byte(secret),
        },
        ContextKey:   "user",
        ErrorHandler: jwtError,
    })
}

func jwtError(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized or invalid token",
    })
}
