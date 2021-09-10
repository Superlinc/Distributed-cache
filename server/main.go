package main

import (
	"Distributed-cache/server/cache"
	"Distributed-cache/server/cluster"
	"Distributed-cache/server/http"
	"Distributed-cache/server/tcp"
	"flag"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	ttl := flag.Int("ttl", 30, "cache time to live")
	cap := flag.Int("capacity", 1000, "cache capacity to live")
	node := flag.String("node", "127.0.0.1", "node address")
	clus := flag.String("cluster", "", "cluster address")
	flag.Parse()
	log.Println("type is", *typ)
	log.Println("ttl is", *ttl)
	log.Println("node is", *node)
	log.Println("cluster is", *clus)
	c := cache.New(*typ, *ttl, *cap)
	n, e := cluster.New(*node, *clus)
	if e != nil {
		panic(e)
	}
	go tcp.New(c, n).Listen()
	http.New(c, n).Listen()
}
