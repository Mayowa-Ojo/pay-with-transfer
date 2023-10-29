package cache

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	*redis.Client
	namespace string
	username  string
	password  string
	tlsConfig *tls.Config
}

type Params func(*RedisClient)

func WithUsername(username string) Params {
	return func(rc *RedisClient) {
		rc.username = username
	}
}
func WithPassword(password string) Params {
	return func(rc *RedisClient) {
		rc.password = password
	}
}
func WithNamespace(namespace string) Params {
	return func(rc *RedisClient) {
		rc.namespace = namespace
	}
}

func New(ctx context.Context, host, port string, params ...Params) Cache {
	address := fmt.Sprintf("%s:%s", host, port)

	rc := &RedisClient{}
	for _, param := range params {
		param(rc)
	}

	c := redis.NewClient(&redis.Options{
		Addr:        address,
		Password:    rc.password,
		Username:    rc.username,
		TLSConfig:   rc.tlsConfig,
		DialTimeout: 15 * time.Second,
		MaxRetries:  10,
	})

	rc.Client = c
	return rc
}

func (c *RedisClient) Ping(ctx context.Context) (interface{}, error) {
	return c.Client.Ping(ctx).Result()
}

func (c *RedisClient) SetNX(ctx context.Context, key string, val interface{}, ttl time.Duration) *redis.BoolCmd {
	return c.Client.SetNX(ctx, fmt.Sprintf("%s:%s", c.namespace, key), val, ttl)
}
