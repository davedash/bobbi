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

func handler(w http.ResponseWriter, r *http.Request, cfg config) {
	route := r.URL.Path[1:]
	if data, ok := cfg.Route[route]; ok {
		fmt.Fprintf(w, "Welcome, %s! want to run %s", route, data.Command)
	} else {
		http.NotFound(w, r)
	}
}

type config struct {
	Main struct {
		Port string
	}

	Route map[string]*struct {
		Command string
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
	for route, data := range cfg.Route {
		fmt.Printf("The route %s will run %s\n", route, data.Command)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, cfg)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
