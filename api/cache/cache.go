package cache

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	CourseSystemRedis *redis.Client
)

func InitRedis(host string, db int, password string) (err error) {
	if CourseSystemRedis != nil {
		return nil
	}
	CourseSystemRedis = redis.NewClient(&redis.Options{Addr: host, DB: db, Password: password})
	if CourseSystemRedis == nil {
		return fmt.Errorf("redis inited is nil")
	}
	return nil
}
