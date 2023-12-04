package utils

import (
	"context"

	"github.com/zuno90/go-ws/configs"
)

func Set(key string, value []byte) error {
	if err := configs.CacheClient.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func Del(key string, value string) {}

func HSet(key, field string, value any) error {
	if err := configs.CacheClient.HSet(context.Background(), key, field, value).Err(); err != nil {
		return err
	}
	return nil
}
