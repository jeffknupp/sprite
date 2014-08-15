package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jeffknupp/sprite/core"
)

func main() {
	log.Println("---Starting Sprite---")
	config := core.ConfigureFromFile("config/sprite.conf")
	serveAt := config.Host + ":" + strconv.Itoa(config.Port)
	http.HandleFunc("/", core.ServeFile)
	log.Println(config.VirtualHosts)
	for key, virtualHost := range config.VirtualHosts {
		log.Println("Adding vhost %s", key)
		go core.ServeVirtualHost(virtualHost)
	}
	http.ListenAndServe(serveAt, nil)
	log.Println("---Stopping Sprite---")
}
