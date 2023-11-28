package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetContacts(c *fiber.Ctx) error {
	contacts, err := GetAllContactsFromRedis(rdb)
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

	if err := SaveContactInRedis(rdb, &contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(contact)
}

func GetContact(c *fiber.Ctx) error {
	id := c.Params("id")

	contact, err := GetContactFromRedis(rdb, id)
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

	if err := UpdateContactInRedis(rdb, id, &contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := DeleteContactFromRedis(rdb, id); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
