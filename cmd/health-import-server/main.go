package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	server "github.com/joeecarter/health-import-server"
	"github.com/joeecarter/health-import-server/storage"
)

var addr string
var metricStores []storage.MetricStore

func init() {
	flag.StringVar(&addr, "addr", ":8080", "the address to start the api on e.g. ':8080'")

	// TODO: Change this to a
	configFilePath, ok := os.LookupEnv("CONFIG_FILE_PATH") // Should there be a default for this?
	if !ok {
		log.Fatalf("Please set the CONFIG_FILE_PATH environment variable.\n")
	}

	var err error
	metricStores, err = storage.LoadMetricStores(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load metric stores: %s.\n", err.Error())
	}

	if len(metricStores) == 0 {
		log.Fatalln("You have zero metric stores configured")
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	http.Handle("/upload", server.NewImportHandler(metricStores))

	log.Printf("Starting web server with addr '%s'...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
