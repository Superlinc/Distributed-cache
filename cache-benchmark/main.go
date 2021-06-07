package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type statistic struct {
	count int
	time  time.Duration
}

type result struct {
	getCount    int
	setCount    int
	missCount   int
	statBuckets []statistic
}

func (r *result) addStatistic(bucket int, stat statistic) {
	if bucket > len(r.statBuckets)-1 {
		newStatBuckets := make([]statistic, bucket+1)
		copy(newStatBuckets, r.statBuckets)
		r.statBuckets = newStatBuckets
	}
	s := r.statBuckets[bucket]
	s.count += stat.count
	s.time += stat.time
	r.statBuckets[bucket] = s
}

func (r *result) addDuration(d time.Duration, typ string) {
	bucket := int(d / time.Millisecond)
	r.addStatistic(bucket, statistic{1, d})
	switch typ {
	case "get":
		r.getCount++
		break
	case "set":
		r.setCount++
		break
	default:
		r.missCount++
		break
	}
}

func (r *result) addResult(src *result) {
	for bucket, stat := range src.statBuckets {
		r.addStatistic(bucket, stat)
	}
	r.getCount += src.getCount
	r.setCount += src.setCount
	r.missCount += src.missCount
}

var typ, server, operation string
var total, valueSize, threads, keyspaceLen, pipeLen int

func init() {
	flag.StringVar(&typ, "type", "redis", "cache server type")
	flag.StringVar(&server, "h", "localhost", "cache server address")
	flag.IntVar(&total, "n", 1000, "total number of request")
	flag.IntVar(&valueSize, "d", 1000, "data size of SET/GET value in bytes")
	flag.IntVar(&threads, "c", 1, "number of parallel connections")
	flag.StringVar(&operation, "t", "set", "test set, could be get/set/mixed")
	flag.IntVar(&keyspaceLen, "r", 0, "keyspace length, use random keys from 0 to keyspaceLen - 1")
	flag.IntVar(&pipeLen, "P", 1, "pipeline length")
	flag.Parse()
	fmt.Println("type is", typ)
	fmt.Println("server is", server)
	fmt.Println("total", total, "request")
	fmt.Println("data size is", valueSize)
	fmt.Println("we have", threads, "connections")
	fmt.Println("operation is", operation)
	fmt.Println("keyspace length is", keyspaceLen)
	fmt.Println("pipeline length is", pipeLen)

	rand.Seed(time.Now().UnixNano())
}

func main() {

}
