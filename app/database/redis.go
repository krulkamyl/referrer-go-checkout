package database

import (
	"context"
	"fmt"
	"referrer/app/config"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client
var CacheChannel chan string

func SetupRedis() {

	db, err := strconv.Atoi(config.Getenv("REDIS_DB"))

	if err != nil {
		panic(err)
	}

	Cache = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Getenv("REDIS_HOST"), config.Getenv("REDIS_PORT")),
		DB:   db,
	})
}

func SetupCacheChannel() {
	CacheChannel = make(chan string)
	go func(ch chan string) {
		for {
			time.Sleep(5 * time.Second)

			key := <-ch

			Cache.Del(context.Background(), key)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key
	}
}
