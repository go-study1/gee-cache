package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	geecache "github.com/go-study1/gee-cache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
}

func main() {
	geecache.NewGroup("hello", 2<<10, geecache.GetterFunc(func(key string) ([]byte, error) {
		if v, ok := db[key]; ok {
			return []byte(v), nil
		}
		return nil, errors.New(fmt.Sprintf("not found key: %s", key))
	}))
	addr := "localhost:9999"
	handle := geecache.NewHttpPool(addr)
	log.Println("geecache servier is runing at", addr)
	log.Fatal(http.ListenAndServe(addr, handle))
}
