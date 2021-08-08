package main

import (
	"Distributed-cache/server/cache"
	"Distributed-cache/server/http"
	"Distributed-cache/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
