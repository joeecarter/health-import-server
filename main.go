package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joeecarter/health-import-server/request"
	"github.com/joeecarter/health-import-server/storage"
)

var addr string
var metricStores map[string]storage.MetricStore

func init() {
	flag.StringVar(&addr, "addr", ":8080", "the address to start the api on e.g. ':8080'")

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

	http.HandleFunc("/", apiHandler(apiExportHandler))

	log.Printf("Starting web server with addr '%s'...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func apiExportHandler(req *http.Request) (string, error) {
	log.Printf("Received request with User-Agent: '%s'\n", req.Header.Get("User-Agent"))

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic("Failed to get body, er r =" + err.Error())
	}

	export, err := request.Parse(b)
	if err != nil {
		return "", err
	}

	populatedMetrics := export.PopulatedMetrics()
	log.Printf("Total metrics: %d (%d populated) Total samples %d\n", len(export.Metrics), len(populatedMetrics), export.TotalSamples())

	// TODO: May want to not fail fast here? run all metric stores before erroring to avoid data loss?
	for metricStoreType, metricStore := range metricStores {
		log.Printf("Starting upload to metric store \"%s\".", metricStoreType)
		if metricStore.Store(populatedMetrics); err != nil {
			log.Printf("Failed upload to metric store \"%s\" with error: %s.", metricStoreType, err.Error())
			return "", err
		}
		log.Printf("Finished upload to metric store \"%s\".", metricStoreType)
	}

	return "Processed request.", nil
}

func apiHandler(handler func(*http.Request) (string, error)) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		msg, err := handler(req)
		if err == nil {
			wr.WriteHeader(200)
			wr.Write([]byte(msg + "\n"))
		} else {
			wr.WriteHeader(500)
			wr.Write([]byte("ERROR: " + err.Error() + "\n"))
		}
	}
}
