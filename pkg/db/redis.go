package db

import "github.com/mushanyux/MSChatServerLib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
