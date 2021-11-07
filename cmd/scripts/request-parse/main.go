package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/joeecarter/health-import-server/request"
)

var shouldLogJson bool

func init() {
	flag.BoolVar(&shouldLogJson, "log", true, "should the parsed json be logged or not?")
}

// Attrmpts to parse the request (from a file) using the request submodule and outputs the result as json
func main() {
	flag.Parse()

	b, err := ioutil.ReadFile("request.json")
	if err != nil {
		panic("Failed to read file request.json err = " + err.Error())
	}

	request.LogUnknownMetrics = true
	req, err := request.Parse(b)
	//_, err = request.Parse(b)
	if err != nil {
		panic("Failed to parse MetricUpload err = " + err.Error())
	}

	if shouldLogJson {
		printJson(req.Metrics)
	}
}

func printJson(v interface{}) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		fmt.Println("Failed to masrhsal to json: ", err.Error())
	}
	fmt.Println(string(b))
}
