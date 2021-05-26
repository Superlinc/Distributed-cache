package http

import (
	"Distributed-cache/server/cache"
	"log"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status", s.statusHandler())
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Println("http start error")
	}
}

func New(c cache.Cache) *Server {
	return &Server{c}
}
