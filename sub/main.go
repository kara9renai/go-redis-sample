package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisCli = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var ctx = context.Background()

func main() {
	subscriber := redisCli.Subscribe(ctx, "send-user-data")

	user := User{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &user); err != nil {
			log.Println(err)
		}

		fmt.Println("Recived message from " + msg.Channel + " channel")
		fmt.Printf("%+v", user)
	}
}
