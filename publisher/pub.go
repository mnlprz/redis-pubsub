package publisher

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func Public(w http.ResponseWriter, r *http.Request) {
	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := redis.Publish("send-user-data", "hola").Err()
	if err != nil {
		log.Panic(err)
	}
	log.Println("message published: ")
}
