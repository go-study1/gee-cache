package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"

	geecache "github.com/go-study1/gee-cache"
)

var db = map[string]string{
	"Tom":  "630",
	"Jack": "589",
	"Sam":  "567",
	"Sam1": "567",
	"Sam2": "567",
	"Sam3": "567",
	"Sam4": "567",
	"Sam5": "567",
}

// 1.创建缓存
// 2.创建缓存http服务，并注册其他缓存节点
// 3.创建api服务

func createGeeCache() *geecache.Group {
	cahce := geecache.NewGroup("hello", 2<<10, geecache.GetterFunc(
		func(key string) ([]byte, error) {
			log.Println("[SlowDb] search key", key)
			if v, ok := db[key]; ok {
				return []byte(v), nil
			}
			return nil, errors.New(fmt.Sprintf("not found key: %s", key))
		}))
	return cahce
}

func startCacheServer(add string, addrs []string, gee *geecache.Group) {
	peers := geecache.NewHttpPool(add)
	peers.Set(addrs...)
	gee.RegisterPeers(peers)
	log.Println("geecache is runing:", add)
	log.Fatal(http.ListenAndServe(add[7:], peers))
}

func startApi(apiAddr string, gee *geecache.Group) {
	http.Handle("/api", http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			key := req.URL.Query().Get("key")
			view, err := gee.Get(key)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			res.Header().Set("Context-Type", "application/octet-stream")
			res.Write(view.ByteSlice())
		}))
	log.Println("api server is runing at", apiAddr)
	log.Fatal(http.ListenAndServe(apiAddr[7:], nil))
}

func main() {
	var port int
	var api bool
	flag.IntVar(&port, "port", 8081, "server port")
	flag.BoolVar(&api, "api", false, "start a api server")
	flag.Parse()

	apiAddr := "http://localhost:9999"
	addMap := map[int]string{
		8081: "http://localhost:8081",
		8082: "http://localhost:8082",
		8083: "http://localhost:8083",
	}
	var addrs []string
	for _, v := range addMap {
		addrs = append(addrs, v)
	}
	gee := createGeeCache()
	if api {
		go startApi(apiAddr, gee)
	}
	startCacheServer(addMap[port], addrs, gee)

}
