package test

import (
	"context"
	"devops-project/src"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
	"testing"
	"time"
)

func TestSaveContactInRedis(t *testing.T) {
	// Redis configuration
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Clear Redis
	defer client.FlushDB(context.Background()).Result()

	redisService := src.NewRedisService(client)

	// Create a contact
	contact := src.Contact{
		Name:  "Test",
		Email: "test@example.com",
		Phone: "1234567890",
	}

	// Try to save the contact
	err := redisService.SaveContactInRedis(&contact)
	if err != nil {
		t.Fatalf("SaveContactInRedis failed: %v", err)
	}

	// Verify that the contact was saved
	savedData, err := client.Get(context.Background(), "contact:"+contact.ID).Result()
	if err != nil {
		t.Fatalf("Failed to get contact from Redis: %v", err)
	}

	var savedContact src.Contact
	err = json.Unmarshal([]byte(savedData), &savedContact)
	if err != nil {
		t.Fatalf("Failed to unmarshal saved contact: %v", err)
	}

	if savedContact.Email != contact.Email || savedContact.Name != contact.Name || savedContact.Phone != contact.Phone {
		t.Error("Saved contact does not match the original contact")
	}

	// Try to save a duplicate contact
	duplicateContact := src.Contact{
		Name:  "Test Duplicate",
		Email: "test@example.com", // même email que le premier contact
		Phone: "0987654321",
	}

	err = redisService.SaveContactInRedis(&duplicateContact)
	if err == nil {
		t.Error("Expected error for duplicate email, got nil")
	}

	// Verify that the duplicate contact was not saved
	_, err = client.Get(context.Background(), "contact:"+strconv.FormatInt(time.Now().UnixNano(), 10)).Result()
	if err != redis.Nil {
		t.Error("Duplicate contact should not have been saved")
	}

	// Delete contact
	err = client.Del(context.Background(), "contact:"+contact.ID).Err()
	if err != nil {
		t.Fatalf("Failed to delete contact: %v", err)
	}

	// Check that the contact was deleted
	_, err = client.Get(context.Background(), "contact:"+contact.ID).Result()
	if err != redis.Nil {
		t.Error("Contact should have been deleted")
	}
}
