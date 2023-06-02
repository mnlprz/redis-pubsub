package subscriber

import (
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func Subscribe(w http.ResponseWriter, r *http.Request) {
	redis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	sub := redis.Subscribe("send-user-data")
	for {
		msg, err := sub.ReceiveMessage()
		if err != nil {
			log.Panic(err)
		}
		log.Println("message subscribed: ", msg.String())
	}
}
