package main

import (
	"log"
	"net/http"

	"github.com/mnlprz/redis-pubsub/publisher"
	"github.com/mnlprz/redis-pubsub/subscriber"
)

func main() {
	http.HandleFunc("/pub", publisher.Public)
	http.HandleFunc("/sub", subscriber.Subscribe)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic(err)
	}
}
