package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ctx = context.Background()

func main() {

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Redis")
	}

	app := fiber.New()

	app.Get("/contacts", GetContacts)
	app.Post("/contacts", CreateContact)
	app.Get("/contacts/:id", GetContact)
	app.Put("/contacts/:id", UpdateContact)
	app.Delete("/contacts/:id", DeleteContact)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal("Fiber stopped: ", err.Error())
		}
	}()

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down Fiber: %v\n", err)
	}
	if err := rdb.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v\n", err)
	}
	<-ctx.Done()
	log.Println("Shutdown successful")
}
