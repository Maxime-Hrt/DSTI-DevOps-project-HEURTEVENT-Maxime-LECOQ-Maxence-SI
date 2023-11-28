//package main
//
//import (
//	"github.com/gofiber/fiber/v2"
//	"github.com/gofiber/fiber/v2/middleware/cors"
//	"github.com/redis/go-redis/v9"
//	"log"
//)
//
//var ctx = context.Background()
//
//func main() {
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//	})
//
//	_, err := rdb.Ping(ctx).Result()
//	if err != nil {
//		log.Fatalf("Error connecting to Redis: %v", err)
//	}
//
//	app := fiber.New()
//
//	app.Use(cors.New(cors.Config{
//		AllowOrigins:     "*",
//		AllowHeaders:     "Origin, Content-Type, Accept",
//		AllowCredentials: true,
//		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",
//	}))
//
//	app.Use(func(c *fiber.Ctx) error {
//		c.Locals("rdb", rdb)
//		return c.Next()
//	})
//
//	app.Get("/contacts", GetContacts)
//	app.Post("/contacts", CreateContact)
//	app.Get("/contacts/:id", GetContact)
//	app.Put("/contacts/:id", UpdateContact)
//	app.Delete("/contacts/:id", DeleteContact)
//
//	log.Fatal(app.Listen(":8080"))
//}

package main

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func main() {
	app := fiber.New()

	// Configurez CORS
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	// Définir les routes
	app.Post("/create", createUser)
	app.Get("/user/:id", readUser) // Route modifiée pour accepter un ID
	app.Put("/update", updateUser)
	app.Delete("/delete", deleteUser)

	log.Fatal(app.Listen(":8080"))
}

func createUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = rdb.Set(ctx, user.ID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("User created successfully")
}

func readUser(c *fiber.Ctx) error {
	// Récupérer l'ID utilisateur depuis l'URL
	userID := c.Params("id")
	//fmt.Println("Requested User ID:", userID) // Ajoutez cette ligne pour le débogage

	val, err := rdb.Get(ctx, userID).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString(val)
}

func updateUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err = rdb.Set(ctx, user.ID, jsonData, 0).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("User updated successfully")
}

func deleteUser(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	_, err := rdb.Del(ctx, user.ID).Result()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendString("User deleted successfully")
}
