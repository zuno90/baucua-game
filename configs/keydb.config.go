package configs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

type Keydb struct {
	host, port string
}

var CacheClient *redis.Client

func ConnectKeydbServer() error {
	conf := &Keydb{host: os.Getenv("KEYDB_HOST"), port: os.Getenv("KEYDB_PORT")}
	addr := fmt.Sprintf("%s:%s", conf.host, conf.port)
	CacheClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := CacheClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Can not ping to Keydb server", err)
	}
	fmt.Println(pong)
	// example
	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for k, v := range session {
		if err := CacheClient.HSet(context.Background(), "user-session:123", k, v).Err(); err != nil {
			panic(err)
		}
	}

	// userSession := CacheClient.HGetAll(context.Background(), "user-session:123").Val()
	// fmt.Println(userSession["name"])

	return nil
}
