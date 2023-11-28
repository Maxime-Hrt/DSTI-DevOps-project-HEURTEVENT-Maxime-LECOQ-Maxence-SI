package src

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var Ctx = context.Background()

func SaveContactInRedis(rdb *redis.Client, contact *Contact) error {
	contact.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	contact.CreatedAt = time.Now()

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	return rdb.Set(Ctx, "contact:"+contact.ID, data, 0).Err()
}

func GetContactFromRedis(rdb *redis.Client, id string) (*Contact, error) {
	val, err := rdb.Get(Ctx, "contact:"+id).Result()
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

func GetAllContactsFromRedis(rdb *redis.Client) ([]Contact, error) {
	keys, err := rdb.Keys(Ctx, "contact:*").Result()
	if err != nil {
		return nil, err
	}

	var contacts []Contact
	for _, key := range keys {
		val, err := rdb.Get(Ctx, key).Result()
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

func DeleteContactFromRedis(rdb *redis.Client, id string) error {
	return rdb.Del(Ctx, "contact:"+id).Err()
}

func UpdateContactInRedis(rdb *redis.Client, id string, contact *Contact) error {
	exists, err := rdb.Exists(Ctx, "contact:"+id).Result()
	if err != nil {
		return err
	}

	if exists == 0 {
		return nil
	}

	data, err := json.Marshal(contact)
	if err != nil {
		return err
	}

	return rdb.Set(Ctx, "contact:"+id, data, 0).Err()
}
