package test

import (
	"bytes"
	"devops-project/src"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"net/http/httptest"
	"testing"
)

func SetupTestApp() *fiber.App {
	app := fiber.New()
	app.Get("/contacts", src.GetContacts)
	app.Post("/contacts", src.CreateContact)
	app.Get("/contacts/:id", src.GetContact)
	app.Get("/contacts/user_email/:email", src.GetContactByEmail)
	app.Put("/contacts/:id", src.UpdateContact)
	app.Delete("/contacts/id/:id", src.DeleteContact)
	app.Delete("/contacts/email/:email", src.DeleteContactByEmail)
	return app
}

func TestGetContacts(t *testing.T) {
	app := fiber.New()
	app.Get("/contacts", src.GetContacts)

	resp, err := app.Test(httptest.NewRequest("GET", "/contacts", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)
}

func TestCreateContact(t *testing.T) {
	app := SetupTestApp()

	contact := src.Contact{
		Name:  "Test",
		Email: "test@example.com",
		Phone: "1234567890",
	}

	// Test valid contact
	contactBytes, _ := json.Marshal(contact)
	reqPost := httptest.NewRequest("POST", "/contacts", bytes.NewReader(contactBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err := app.Test(reqPost)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 201, respPost.StatusCode)

	// Test invalid contact (mail duplicate)
	contactBytes, _ = json.Marshal(contact)
	reqPost = httptest.NewRequest("POST", "/contacts", bytes.NewReader(contactBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err = app.Test(reqPost)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 500, respPost.StatusCode)

	// Delete contact
	contactPath := fmt.Sprintf("/contacts/email/%s", contact.Email)
	reqDelete := httptest.NewRequest("DELETE", contactPath, nil)
	respDelete, err := app.Test(reqDelete)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 204, respDelete.StatusCode)
}

func TestGetContact(t *testing.T) {
	app := SetupTestApp()

	contact := src.Contact{
		Name:  "Test",
		Email: "test@example.com",
		Phone: "1234567890",
	}

	// Create contact
	contactBytes, _ := json.Marshal(contact)
	reqPost := httptest.NewRequest("POST", "/contacts", bytes.NewReader(contactBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err := app.Test(reqPost)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 201, respPost.StatusCode)

	// Get contact
	contactPath := fmt.Sprintf("/contacts/user_email/%s", contact.Email)
	reqGet := httptest.NewRequest("GET", contactPath, nil)
	respGet, err := app.Test(reqGet)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, respGet.StatusCode)

	// Delete contact
	contactPath = fmt.Sprintf("/contacts/email/%s", contact.Email)
	reqDelete := httptest.NewRequest("DELETE", contactPath, nil)
	respDelete, err := app.Test(reqDelete)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 204, respDelete.StatusCode)

}

func TestHealthCheck(t *testing.T) {
	app := fiber.New()
	app.Get("/health", src.HealthCheck)

	resp, err := app.Test(httptest.NewRequest("GET", "/health", nil))

	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)
}

func TestEmailDuplicate(t *testing.T) {
	app := SetupTestApp()

	contact := src.Contact{
		Name:  "Test",
		Email: "test@example.com",
		Phone: "1234567890",
	}

	// Create contact
	contactBytes, _ := json.Marshal(contact)
	reqPost := httptest.NewRequest("POST", "/contacts", bytes.NewReader(contactBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err := app.Test(reqPost)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 201, respPost.StatusCode)

	// Create contact with same email
	contactBytes, _ = json.Marshal(contact)
	reqPost = httptest.NewRequest("POST", "/contacts", bytes.NewReader(contactBytes))
	reqPost.Header.Set("Content-Type", "application/json")
	respPost, err = app.Test(reqPost)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 500, respPost.StatusCode)

	// Delete contact
	contactPath := fmt.Sprintf("/contacts/email/%s", contact.Email)
	reqDelete := httptest.NewRequest("DELETE", contactPath, nil)
	respDelete, err := app.Test(reqDelete)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 204, respDelete.StatusCode)
}
