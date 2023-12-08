package main

import (
	"context"
	_ "devops-project/docs"
	"devops-project/src"
	"github.com/ansrivas/fiberprometheus/v2"
	fiberSwagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_, err := src.Rdb.Ping(src.Ctx).Result()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Redis")
	}

	app := fiber.New()

	// CORS middleware configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	prometheus := fiberprometheus.New("app")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	app.Static("/", "./src/public")

	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	app.Get("/health", src.HealthCheck)
	app.Get("/version", src.Version)

	app.Get("/contacts", src.GetContacts)
	app.Post("/contacts", src.CreateContact)
	app.Get("/contacts/:id", src.GetContact)
	app.Get("/contacts/user_email/:email", src.GetContactByEmail)
	app.Put("/contacts/:id", src.UpdateContact)
	app.Delete("/contacts/id/:id", src.DeleteContact)
	app.Delete("/contacts/email/:email", src.DeleteContactByEmail)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Fatal("Fiber stopped: ", err.Error())
		}
	}()

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Printf("Error shutting down Fiber: %v\n", err)
	}
	if err := src.Rdb.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v\n", err)
	}
	<-ctx.Done()
	log.Println("Shutdown successful")
}
