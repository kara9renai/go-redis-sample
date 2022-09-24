package main

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var redisCli = redis.NewClient(&redis.Options{
	Addr: "redis-sample:6379",
})

func main() {
	const defaultPort = ":8080"

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello there")
	})
	app.Listen(defaultPort)
}
