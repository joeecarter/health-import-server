package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":8080", "the address to start the api on e.g. ':8080'")
}

// Tiny utility that dumps the incomming json request to a file to aid development of this tool.
func main() {
	flag.Parse()

	http.HandleFunc("/", uploadHandler)

	log.Printf("Starting web server with addr '%s'...\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("Received request with User-Agent: '%s'\n", r.Header.Get("User-Agent"))

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("Failed to get body, er r =" + err.Error())
	}

	writeRequestFile(b)

	w.WriteHeader(200)
	w.Write([]byte("Written to request.json"))
	log.Printf("Written to request.json")
}

func writeRequestFile(b []byte) {
	err := ioutil.WriteFile("request.json", b, 0600)
	if err != nil {
		fmt.Printf("Failed to write $CWD/request.json err = '%s'.\n", err)
	}
}
