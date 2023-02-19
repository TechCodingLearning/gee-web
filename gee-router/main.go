package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello Gee</h1>
(2)
$ curl "http://localhost:9999/hello?name=geektutu"
hello geektutu, you're at /hello
(3)
$ curl "http://localhost:9999/hello/geektutu"
hello geektutu, you're at /hello/geektutu
(4)
$ curl "http://localhost:9999/assets/css/geektutu.css"
{"filepath":"css/geektutu.css"}
(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"gee-router/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(ctx *gee.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(ctx *gee.Context) {
		// expectxt /hello>name=geektutu
		ctx.String(http.StatusOK, "helllo %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *gee.Context) {
		// expect /hello/geektutu
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Params["name"], ctx.Path)
	})

	r.GET("/assets/*filepath", func(ctx *gee.Context) {
		ctx.JSON(http.StatusOK, gee.H{"filepath": ctx.Params["filepath"]})
	})

	r.Run(":9999")
}
