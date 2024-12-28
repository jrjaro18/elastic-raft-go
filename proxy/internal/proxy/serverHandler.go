package proxy

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

// serverStatus is the handler for the base server route.
func (p *Proxy) serverStatus(c *fiber.Ctx) error {
	return c.SendString("Proxy Working")
}

// registerServerHandler handles server registration requests.
func (p *Proxy) registerServerHandler(c *fiber.Ctx) error {
	// Parse the request body
	body := new(struct {
		Addr string `json:"addr"`
	})
	if err := c.BodyParser(body); err != nil {
		log.Println("Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validate the address
	if body.Addr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Address field is required",
		})
	}

	// Register the server
	p.registerServer(body.Addr)

	// Respond with success and the updated server list
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Registration completed successfully",
		"serverList": p.GetListedServer([]string{body.Addr}),
	})
}
