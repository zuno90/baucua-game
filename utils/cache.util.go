package utils

import (
	"context"

	"github.com/zuno90/go-ws/configs"
)

func Get(key string) (interface{}, error) {
	value, err := configs.CacheClient.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func Set(key string, value []byte) error {
	if err := configs.CacheClient.Set(context.Background(), key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}

func Del(key string) error {
	if err := configs.CacheClient.Del(context.Background(), key).Err(); err != nil {
		return err
	}
	return nil
}

func Clear() error {
	keys, err := configs.CacheClient.Keys(context.Background(), "*").Result()
	if err != nil {
		return err
	}
	if len(keys) > 0 {
		if err := configs.CacheClient.Del(context.Background(), keys...).Err(); err != nil {
			return err
		}
	}
	return nil
}

func HSet(key, field string, value any) error {
	if err := configs.CacheClient.HSet(context.Background(), key, field, value).Err(); err != nil {
		return err
	}
	return nil
}
