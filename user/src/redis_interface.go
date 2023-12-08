package src

import (
	"github.com/redis/go-redis/v9"
	"os"
)

type RedisClient interface {
	SaveContactInRedis(rdb *redis.Client, contact *Contact) error
	GetContactFromRedis(rdb *redis.Client, id string) (*Contact, error)
	GetAllContactsFromRedis(rdb *redis.Client) ([]Contact, error)
	GetContactIdFromEmail(rdb *redis.Client, email string) (string, error)
	DeleteContactFromRedis(rdb *redis.Client, id string) error
	DeleteContactByEmail(rdb *redis.Client, email string) error
	UpdateContactInRedis(rdb *redis.Client, id string, updatedContact *Contact) error
}

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var redisHost = getEnvWithDefault("REDIS_HOST", "localhost")
var redisPort = getEnvWithDefault("REDIS_PORT", "6379")
var redisPassword = getEnvWithDefault("REDIS_PASSWORD", "")

var Rdb = redis.NewClient(&redis.Options{
	Addr:     redisHost + ":" + redisPort,
	Password: redisPassword,
	DB:       0,
})

//var Rdb = redis.NewClient(&redis.Options{
//	Addr:     "redis:6379",
//	Password: "",
//	DB:       0,
//})
