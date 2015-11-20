// Webserver
package main

import (
	"flag"
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
	"net/http"
)

var configFile = flag.String("config", "/etc/bobbi/bobbi.conf", "INI file for bobbi")

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

type config struct {
	Main struct {
		Port string
	}
	Routes struct {
	}
}

func getConfig(filename string) config {
	cfg := new(config)

	err := gcfg.ReadFileInto(cfg, filename)

	if err != nil {
		log.Fatalf("Failed reading config: %s", err)
	}

	return *cfg
}

func main() {
	flag.Parse()
	cfg := getConfig(*configFile)
	port := cfg.Main.Port
	if port == "" {
		port = "8000"
	}

	fmt.Println("Listening on port:", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
