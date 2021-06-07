package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type httpClient struct {
	*http.Client
	server string
}

func (c *httpClient) get(key string) string {
	resp, e := c.Get(c.server + key)
	if e != nil {
		log.Println(key)
		panic(e)
	}
	if resp.StatusCode == http.StatusNotFound {
		return ""
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		panic(e)
	}
	return string(body)
}

func (c *httpClient) set(key, value string) {
	req, e := http.NewRequest(http.MethodPut, c.server+key, strings.NewReader(value))
	if e != nil {
		log.Println(key)
		panic(e)
	}
	resp, e := c.Do(req)
	if e != nil {
		log.Println(key)
		panic(e)
	}
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
}

func (c *httpClient) Run(cmd *Cmd) {
	switch cmd.Name {
	case "get":
		cmd.Value = c.get(cmd.Key)
		return
	case "set":
		c.set(cmd.Key, cmd.Value)
		return
	default:
		panic("unknown cmd name " + cmd.Name)
	}
}

func (c *httpClient) PipelinedRun([]*Cmd) {
	panic("httpClient pipelined run not implement")
}

func newHttpClient(server string) *httpClient {
	client := &http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: 1}}
	return &httpClient{client, "http://" + server + ":12345/cache/"}
}
