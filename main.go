package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jeffknupp/sprite/core"
)

type Sprite struct {
	Config *core.Configuration
}

func New() *Sprite {
	return &Sprite{Config: core.ConfigureFromFile("config/sprite.conf")}
}

func (s *Sprite) Run() {
	serveAt := s.Config.Host + ":" + strconv.Itoa(s.Config.Port)
	http.HandleFunc("/", core.ServeFile)
	for key, virtualHost := range s.Config.VirtualHosts {
		log.Println("Adding vhost %s", key)
		go core.ServeVirtualHost(virtualHost)
	}
	http.ListenAndServe(serveAt, nil)
}

func startServer() {
	s := New()
	s.Run()
}

func main() {
	log.Println("---Starting Sprite---")
	startServer()
	log.Println("---Stopping Sprite---")
}
