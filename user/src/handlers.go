package src

import (
	"github.com/gofiber/fiber/v2"
)

var redisService = NewRedisService(Rdb)

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func Index(c *fiber.Ctx) error {
	return c.SendString("Hello Azure 🙏🏻!")
}

// GetContacts godoc
// @Summary Get list of contacts
// @Description Retrieve a list of all contacts
// @Tags contacts
// @Accept json
// @Produce json
// @Success 200 {array} Contact
// @Router /contacts [get]
func GetContacts(c *fiber.Ctx) error {
	contacts, err := redisService.GetAllContactsFromRedis()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(contacts)
}

// CreateContact godoc
// @Summary Create a new contact
// @Description Add a new contact to the database
// @Tags contacts
// @Accept json
// @Produce json
// @Param contact body Contact true "Contact to add"
// @Success 201 {object} Contact "Contact successfully created"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /contacts [post]
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

// GetContact godoc
// @Summary Get a contact by ID
// @Description Retrieve a contact by its unique ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Success 200 {object} Contact "Contact found"
// @Failure 404 {string} string "Contact not found"
// @Router /contacts/{id} [get]
func GetContact(c *fiber.Ctx) error {
	id := c.Params("id")

	contact, err := redisService.GetContactFromRedis(id)
	if err != nil {
		return c.Status(404).SendString(err.Error())
	}

	return c.JSON(contact)
}

// UpdateContact godoc
// @Summary Update a contact
// @Description Update an existing contact by its unique ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Param contact body Contact true "Updated contact information"
// @Success 200 {object} Contact "Contact successfully updated"
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /contacts/{id} [put]
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

// DeleteContact godoc
// @Summary Delete a contact by ID
// @Description Remove a contact from the database by its unique ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "Contact ID"
// @Success 204 "Contact successfully deleted"
// @Failure 500 {string} string "Internal server error"
// @Router /contacts/id/{id} [delete]
func DeleteContact(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := redisService.DeleteContactFromRedis(id); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}

// DeleteContactByEmail godoc
// @Summary Delete a contact by email
// @Description Remove a contact from the database by email
// @Tags contacts
// @Accept json
// @Produce json
// @Param email path string true "Contact Email"
// @Success 204 "Contact successfully deleted"
// @Failure 500 {string} string "Internal server error"
// @Router /contacts/email/{email} [delete]
func DeleteContactByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	if err := redisService.DeleteContactByEmail(email); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.SendStatus(204)
}
