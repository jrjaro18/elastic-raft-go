package proxy

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func (p *Proxy) proxyStatus(c *fiber.Ctx) error {
	return c.SendString("Proxy Working")
}

func (p *Proxy) registerServerHandler(c *fiber.Ctx) error {
	body := new(struct {
		Addr string `json:"addr"`
	})

	if err := c.BodyParser(body); err != nil {
		log.Println("Failed to parse request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if body.Addr == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Address field is required",
		})
	}

	p.addNewServer(body.Addr)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Registration completed successfully",
		"serverList": p.GetListedServer([]string{body.Addr}),
	})
}
