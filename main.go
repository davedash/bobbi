// Webserver
package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

type config struct {
	Main struct {
		Port string
	}
}

func getConfig() config {
	cfg := new(config)

	err := gcfg.ReadFileInto(cfg, "./config-example.conf")

	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
		os.Exit(1)
	}

	return *cfg
}

func main() {
	cfg := getConfig()
	port := cfg.Main.Port
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
