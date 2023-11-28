package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func GetContacts(c *fiber.Ctx) error {
	rdb := c.Locals("rdb").(*redis.Client)
	ctx := c.Context()
	contacts, err := GetAllContactsFromRedis(ctx, rdb)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(contacts)
}

func CreateContact(c *fiber.Ctx) error {
	rdb := c.Locals("rdb").(*redis.Client)

	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	ctx := c.Context()
	if err := SaveContactInRedis(ctx, rdb, &contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(contact)
}

func GetContact(c *fiber.Ctx) error {
	rdb := c.Locals("rdb").(*redis.Client)
	id := c.Params("id")
	ctx := c.Context()

	contact, err := GetContactFromRedis(ctx, rdb, id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.JSON(contact)
}

func UpdateContact(c *fiber.Ctx) error {
	rdb := c.Locals("rdb").(*redis.Client)
	id := c.Params("id")
	ctx := c.Context()

	var contact Contact
	if err := c.BodyParser(&contact); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if err := UpdateContactInRedis(ctx, rdb, id, &contact); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(contact)
}

func DeleteContact(c *fiber.Ctx) error {
	rdb := c.Locals("rdb").(*redis.Client)
	id := c.Params("id")
	ctx := c.Context()

	if err := DeleteContactFromRedis(ctx, rdb, id); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
