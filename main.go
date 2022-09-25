package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var redisCli = redis.NewClient(&redis.Options{
	Addr: "redis-sample:6379",
})

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	const defaultPort = ":8080"

	fmt.Print("start\n")
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello there")
	})
	app.Post("/", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			log.Println(err)
		}
		payload, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
		}
		if err := redisCli.Publish(ctx, "send-user-data", payload).Err(); err != nil {
			log.Println(err)
		}

		return c.SendStatus(200)
	})
	app.Listen(defaultPort)
}
