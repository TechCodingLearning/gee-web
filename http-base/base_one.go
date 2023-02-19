package main

//$ curl http:/localhost:9999/
//URL.Path = "/"
//
//$ curl http://localhost:9999/hello
//Header["User-Agent"] = ["curl/7.82.0"]
//Header["Accept"] = ["*/*"]

import (
	"fmt"
	"log"
	"net/http"
)

func TestOne() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
