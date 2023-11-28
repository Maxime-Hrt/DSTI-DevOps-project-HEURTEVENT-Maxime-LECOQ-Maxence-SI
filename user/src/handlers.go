package src

import (
	"github.com/gofiber/fiber/v2"
)

var redisService = NewRedisService(Rdb)

func GetContacts(c *fiber.Ctx) error {
	contacts, err := redisService.GetAllContactsFromRedis()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(contacts)
}

func CreateContact(c *fiber.Ctx) error {
	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := redisService.SaveContactInRedis(&contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(contact)
}

func GetContact(c *fiber.Ctx) error {
	id := c.Params("id")

	contact, err := redisService.GetContactFromRedis(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.JSON(contact)
}

func UpdateContact(c *fiber.Ctx) error {
	id := c.Params("id")

	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := redisService.UpdateContactInRedis(id, &contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := redisService.DeleteContactFromRedis(id); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
