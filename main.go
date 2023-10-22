package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "success", "message": "hello world, i'm ok!", "data": nil})
	})

	log.Fatal(app.Listen(":8080"))
}
