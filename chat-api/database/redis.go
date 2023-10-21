package database

import (
	"chatjobsity/env"
	"fmt"

	"github.com/go-redis/redis"
)

type Redis struct {
	env env.EnvApp
}

func NewRedis(env env.EnvApp) *Redis {
	return &Redis{env: env}
}

func (m *Redis) Start() (*redis.Client, error) {
	opt, err := redis.ParseURL(fmt.Sprintf("redis://%s:%s", m.env.RedisHost, m.env.RedisPort))
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	return client, nil
}
