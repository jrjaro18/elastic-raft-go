package middleware

import "github.com/gofiber/fiber/v2"

func ApiKeyMiddleware(requiredKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		if apiKey == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "API key is missing",
			})
		}
		if apiKey != requiredKey {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Invalid API key",
			})
		}
		return c.Next()
	}
}
