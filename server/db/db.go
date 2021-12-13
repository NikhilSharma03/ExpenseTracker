package db

import "github.com/go-redis/redis"

var rdb *redis.Client

func ConnectDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return rdb
}
