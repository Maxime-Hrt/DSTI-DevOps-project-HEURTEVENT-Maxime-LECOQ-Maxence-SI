package src

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type RedisService struct {
	Client *redis.Client
}

var Ctx = context.Background()

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		Client: client,
	}
}

func (r *RedisService) SaveContactInRedis(contact *Contact) error {
	iter := r.Client.Scan(Ctx, 0, "contact:*", 0).Iterator()
	for iter.Next(Ctx) {
		existingContactData, err := r.Client.Get(Ctx, iter.Val()).Result()
		if err != nil {
			return err
		}

		var existingContact Contact
		err = json.Unmarshal([]byte(existingContactData), &existingContact)
		if err != nil {
			return err
		}
		if existingContact.Email == contact.Email {
			return fmt.Errorf("this email already exists")
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}

	contact.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	contact.CreatedAt = time.Now()

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	return r.Client.Set(Ctx, "contact:"+contact.ID, data, 0).Err()
}

func (r *RedisService) GetContactFromRedis(id string) (*Contact, error) {
	val, err := r.Client.Get(Ctx, "contact:"+id).Result()
	if err != nil {
		return nil, err
	}

	var contact Contact
	err = json.Unmarshal([]byte(val), &contact)
	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (r *RedisService) GetContactIdFromEmail(email string) (string, error) {
	var contacts []Contact
	keys, err := r.Client.Keys(Ctx, "contact:*").Result()
	if err != nil {
		return "", err
	}

	for _, key := range keys {
		val, err := r.Client.Get(Ctx, key).Result()
		if err != nil {
			return "", err
		}

		var contact Contact
		err = json.Unmarshal([]byte(val), &contact)
		if err != nil {
			return "", err
		}

		contacts = append(contacts, contact)
	}

	var contactId string
	for _, contact := range contacts {
		if contact.Email == email {
			contactId = contact.ID
		}
	}

	if contactId == "" {
		return "", fmt.Errorf("contact not found")
	}

	return contactId, nil
}

func (r *RedisService) GetAllContactsFromRedis() ([]Contact, error) {
	keys, err := r.Client.Keys(Ctx, "contact:*").Result()
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	for _, key := range keys {
		val, err := r.Client.Get(Ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var contact Contact
		err = json.Unmarshal([]byte(val), &contact)
		if err != nil {
			return nil, err
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *RedisService) DeleteContactFromRedis(id string) error {
	key := "contact:" + id

	exists, err := r.Client.Exists(Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	if exists == 0 {
		return nil
	}

	return r.Client.Del(Ctx, "contact:"+id).Err()
}

func (r *RedisService) DeleteContactByEmail(email string) error {
	var contacts []Contact
	keys, err := r.Client.Keys(Ctx, "contact:*").Result()
	if err != nil {
		return err
	}

	for _, key := range keys {
		val, err := r.Client.Get(Ctx, key).Result()
		if err != nil {
			return err
		}

		var contact Contact
		err = json.Unmarshal([]byte(val), &contact)
		if err != nil {
			return err
		}

		contacts = append(contacts, contact)
	}

	var emailKey string
	for _, contact := range contacts {
		if contact.Email == email {
			emailKey = "contact:" + contact.ID
		}
	}

	if emailKey == "" {
		return nil
	}

	return r.Client.Del(Ctx, emailKey).Err()
}

func (r *RedisService) UpdateContactInRedis(id string, updatedContact *Contact) error {
	key := "contact:" + id

	exists, err := r.Client.Exists(Ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}

	if exists == 0 {
		return nil
	}

	existingData, err := r.Client.Get(Ctx, key).Result()
	if err != nil {
		return err
	}

	var existingContact Contact
	err = json.Unmarshal([]byte(existingData), &existingContact)
	if err != nil {
		return err
	}

	if updatedContact.Name != "" {
		existingContact.Name = updatedContact.Name
	}
	if updatedContact.Email != "" {
		existingContact.Email = updatedContact.Email
	}
	if updatedContact.Phone != "" {
		existingContact.Phone = updatedContact.Phone
	}

	data, err := json.Marshal(existingContact)
	if err != nil {
		return err
	}

	return r.Client.Set(Ctx, key, data, 0).Err()
}
