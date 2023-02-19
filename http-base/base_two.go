package main

import (
	"fmt"
	"log"
	"net/http"
)

//$ curl http:/localhost:9999/
//URL.PATH = "/"
//$ curl http://localhost:9999/hello
//Header["User-Agent"]
//Header["User-Agent"] = ["curl/7.82.0"]
//Header["Accept"] = ["*/*"]
//$ curl http:/localhost:9999/world
//404 NOT FOUND: /world

type Engine struct {
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.PATH = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func TestTwo() {
	e := new(Engine)
	log.Fatalln(http.ListenAndServe(":9999", e))
}
