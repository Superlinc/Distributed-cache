package main

import (
	"Distributed-cache/server/cache"
	"Distributed-cache/server/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
