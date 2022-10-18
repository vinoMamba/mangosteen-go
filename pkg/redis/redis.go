package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

var once sync.Once

var Redis *RedisClient

func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address string, username string, password string, db int) *RedisClient {
	rdc := &RedisClient{}
	rdc.Context = context.Background()
	rdc.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rdc.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	return rdc
}

func (r *RedisClient) Ping() error {
	_, err := r.Client.Ping(r.Context).Result()
	return err
}
func (r *RedisClient) Get(key string) string {
	result, err := r.Client.Get(r.Context, key).Result()
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return result
}
func (r *RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := r.Client.Set(r.Context, key, value, expiration).Err(); err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
