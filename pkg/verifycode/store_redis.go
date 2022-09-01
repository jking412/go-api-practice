package verifycode

import (
	"go-api-practice/pkg/app"
	"go-api-practice/pkg/config"
	"go-api-practice/pkg/redis"
	"time"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	keyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug.expire_time"))
	}

	return s.RedisClient.Set(s.keyPrefix+key, value, ExpireTime)
}

func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.keyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key string, answer string, clear bool) bool {
	value := s.Get(key, clear)
	return value == answer
}
