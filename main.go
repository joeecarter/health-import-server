package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Printf("URL: %s\n", r.URL)
		//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	
		b, err := ioutil.ReadAll(r.Body)
		if (err == nil) {
			fmt.Println(string(b))
		} else {
			fmt.Println("Error on printing body:", err)
		}
	
		w.WriteHeader(200)
	})

	http.ListenAndServe(":8001", nil)
}
