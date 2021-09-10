package cache

import "log"

func New(typ string, ttl, cap int) Cache {
	var c Cache
	if typ == "inmemory" {
		c = newInMemoryCache(ttl, cap)
	}
	if c == nil {
		panic("unknown cache type " + typ)
	}
	log.Println(typ, "ready to serve")
	return c
}
