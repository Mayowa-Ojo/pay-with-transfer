package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:generate mockgen -source=cache.go -destination cache_mock.go -package cache . Cache

type Cache interface {
	Ping(context.Context) (interface{}, error)
	SetNX(context.Context, string, interface{}, time.Duration) *redis.BoolCmd
}
