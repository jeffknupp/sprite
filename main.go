package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

type VirtualHost struct {
	Host string
	Port int
}

type Configuration struct {
	DocumentRoot string
	Host         string
	Port         int
	IndexFile    bool
	VirtualHosts map[string]VirtualHost
}

func NewConfiguration() *Configuration {
	return &Configuration{IndexFile: true, VirtualHosts: make(map[string]VirtualHost)}
}

var config = *NewConfiguration()

func configureFromFile(path string) *Configuration {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	configData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := toml.Decode(string(configData), &config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if config.IndexFile && path == "/" {
		path = "/index.html"
	}
	path = config.DocumentRoot + path
	log.Println(path)
	http.ServeFile(w, r, path)
}

func serveVirtualHost(vhost VirtualHost) {
	serveAt := vhost.Host + ":" + strconv.Itoa(vhost.Port)
	log.Println("Serving VirtualHost at [" + serveAt + "]")
	http.ListenAndServe(serveAt, nil)
}
func main() {
	log.Println("---Starting Sprite---")
	config := configureFromFile("config/sprite.conf")
	serveAt := config.Host + ":" + strconv.Itoa(config.Port)
	http.HandleFunc("/", serveFile)
	log.Println(config.VirtualHosts)
	for key, virtualHost := range config.VirtualHosts {
		log.Println("Adding vhost %s", key)
		go serveVirtualHost(virtualHost)
	}
	http.ListenAndServe(serveAt, nil)
	log.Println("---Stopping Sprite---")
}
