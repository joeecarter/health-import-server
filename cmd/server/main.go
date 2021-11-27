package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	server "github.com/joeecarter/health-import-server"
)

var Version = "0.0.0"

var addr string
var configFilePath string

var metricStores []server.MetricStore

func init() {
	flag.StringVar(&addr, "addr", ":8080", "The address to start the server on e.g. ':8080'")
	flag.StringVar(&configFilePath, "config", "", "Path to the config file (optional).")
	flag.Parse()

	var err error
	metricStores, err = server.LoadMetricStores(configFilePath)
	if err != nil {
		fmt.Printf("Failed to load metric stores: %s.\n", err.Error())
		os.Exit(1)
	}

	if len(metricStores) == 0 {
		printConfigurationExplanation()
		os.Exit(1)
	}
}

func main() {

	http.Handle("/upload", server.NewImportHandler(metricStores))

	log.Printf("Starting health-import-server v%s with addr '%s'...\n", Version, addr)
	log.Printf("Point Auto Export to /upload")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func printConfigurationExplanation() {
	fmt.Printf("You have no metric stores configured.\n\n")

	fmt.Printf("For configuration you have two options:\n")
	fmt.Printf("1. Environment variables\n")
	fmt.Printf("2. Config file\n")
	fmt.Printf("\n")

	fmt.Println("For option 1 you can configure an influxdb by setting these environment variables:")
	fmt.Println("- INFLUX_HOSTNAME")
	fmt.Println("- INFLUX_BUCKET")
	fmt.Println("- INFLUX_TOKEN")
	fmt.Println("- INFLUX_ORG")
	fmt.Printf("\n")

	fmt.Println("For option 2 you can set a config file with the --config flag:")
	fmt.Println("[")
	fmt.Println("\t{")
	fmt.Println("\t\t\"type\": \"influxdb\",")
	fmt.Println("\t\t\"hostname\": \"<your hostname here>\",")
	fmt.Println("\t\t\"token\": \"<your token here>\",")
	fmt.Println("\t\t\"org\": \"<your org here>\",")
	fmt.Println("\t\t\"bucket\": \"<your bucket here>\",")
	fmt.Println("\t}")
	fmt.Println("]")
}
