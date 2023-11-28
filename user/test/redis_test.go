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
	// Configuration du client Redis pour les tests
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // ou l'adresse de votre instance de test Redis
		Password: "",               // aucun mot de passe pour Redis
		DB:       0,                // Utilisez une base de données distincte pour les tests
	})

	// Assurez-vous de nettoyer la base de données de test après le test
	defer client.FlushDB(context.Background()).Result()

	redisService := src.NewRedisService(client)

	// Créer un contact
	contact := src.Contact{
		Name:  "Test",
		Email: "test@example.com",
		Phone: "1234567890",
	}

	// Tester la sauvegarde du contact
	err := redisService.SaveContactInRedis(&contact)
	if err != nil {
		t.Fatalf("SaveContactInRedis failed: %v", err)
	}

	// Vérifier que le contact est bien enregistré
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

	// Tester la sauvegarde d'un contact avec un email en double
	duplicateContact := src.Contact{
		Name:  "Test Duplicate",
		Email: "test@example.com", // même email que le premier contact
		Phone: "0987654321",
	}

	err = redisService.SaveContactInRedis(&duplicateContact)
	if err == nil {
		t.Error("Expected error for duplicate email, got nil")
	}

	// Vérifier que l'email en double n'a pas été enregistré
	_, err = client.Get(context.Background(), "contact:"+strconv.FormatInt(time.Now().UnixNano(), 10)).Result()
	if err != redis.Nil {
		t.Error("Duplicate contact should not have been saved")
	}

	// Supprimer le contact
	err = client.Del(context.Background(), "contact:"+contact.ID).Err()
	if err != nil {
		t.Fatalf("Failed to delete contact: %v", err)
	}

	// Vérifier que le contact a bien été supprimé
	_, err = client.Get(context.Background(), "contact:"+contact.ID).Result()
	if err != redis.Nil {
		t.Error("Contact should have been deleted")
	}
}
