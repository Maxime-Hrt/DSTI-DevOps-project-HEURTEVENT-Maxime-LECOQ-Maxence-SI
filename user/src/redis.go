package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func SaveContactInRedis(ctx context.Context, rdb *redis.Client, contact *Contact) error {
	contact.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	contact.CreatedAt = time.Now()

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, "contact:"+contact.ID, data, 0).Err()
}

func GetContactFromRedis(ctx context.Context, rdb *redis.Client, id string) (*Contact, error) {
	val, err := rdb.Get(ctx, "contact:"+id).Result()
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

func GetAllContactsFromRedis(ctx context.Context, rdb *redis.Client) ([]Contact, error) {
	keys, err := rdb.Keys(ctx, "contact:*").Result()
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	for _, key := range keys {
		val, err := rdb.Get(ctx, key).Result()
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

func DeleteContactFromRedis(ctx context.Context, rdb *redis.Client, id string) error {
	return rdb.Del(ctx, "contact:"+id).Err()
}

func UpdateContactInRedis(ctx context.Context, rdb *redis.Client, id string, contact *Contact) error {
	exists, err := rdb.Exists(ctx, "contact:"+id).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return nil // ou retourner une erreur si vous préférez
	}

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, "contact:"+id, data, 0).Err()
}
