package src

import "github.com/redis/go-redis/v9"

type RedisClient interface {
	SaveContactInRedis(rdb *redis.Client, contact *Contact) error
	GetContactFromRedis(rdb *redis.Client, id string) (*Contact, error)
	GetAllContactsFromRedis(rdb *redis.Client) ([]Contact, error)
	DeleteContactFromRedis(rdb *redis.Client, id string) error
	DeleteContactByEmail(rdb *redis.Client, email string) error
	UpdateContactInRedis(rdb *redis.Client, id string, updatedContact *Contact) error
}

var Rdb = redis.NewClient(&redis.Options{
	Addr:     "redis:6379",
	Password: "",
	DB:       0,
})
