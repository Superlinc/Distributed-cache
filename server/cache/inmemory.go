package cache

import (
	"container/list"
	"sync"
	"time"
)

type entry struct {
	k       string
	v       []byte
	created time.Time
}

type inMemoryCache struct {
	cache map[string]*list.Element
	l     *list.List
	cap   int
	mutex sync.RWMutex
	Stat
	ttl time.Duration
}

func (c *inMemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	e, ok := c.cache[k]
	c.add(k, v)
	if ok {
		e.Value = &entry{k, v, time.Now()}
		c.l.MoveToFront(e)
	} else {
		c.l.PushFront(&list.Element{
			Value: &entry{k, v, time.Now()},
		})
		if c.l.Len() > c.cap {
			ele := c.l.Back()
			c.l.Remove(c.l.Back())
			ety := ele.Value.(*entry)
			delete(c.cache, ety.k)
			c.del(ety.k, ety.v)
		}
	}
	return nil
}

func (c *inMemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	ele := c.cache[k]
	c.l.MoveToFront(ele)
	return c.cache[k].Value.(*entry).v, nil
}

func (c *inMemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	ele, exist := c.cache[k]
	if exist {
		delete(c.cache, ele.Value.(*entry).k)
		c.del(k, ele.Value.(*entry).v)
	}
	return nil
}

func (c *inMemoryCache) GetStat() Stat {
	return c.Stat
}

func newInMemoryCache(ttl, cap int) *inMemoryCache {
	c := &inMemoryCache{
		make(map[string]*list.Element),
		list.New(),
		cap,
		sync.RWMutex{},
		Stat{},
		time.Duration(ttl) * time.Second}
	if ttl > 0 {
		go c.expirer()
	}
	return c
}

func (c *inMemoryCache) expirer() {
	for {
		time.Sleep(c.ttl)
		c.mutex.RLock()
		for k, ele := range c.cache {
			c.mutex.RUnlock()
			ety := ele.Value.(*entry)
			if ety.created.Add(c.ttl).Before(time.Now()) {
				c.Del(k)
			}
			c.mutex.RLock()
		}
		c.mutex.RUnlock()
	}
}
