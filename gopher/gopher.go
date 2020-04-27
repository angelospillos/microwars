package main

import (
	"github.com/valyala/fasthttp"
	"time"
)

var contentTypeJson = []byte("application/json")

func handlerFunc(ctx *fasthttp.RequestCtx) {
	resp := &ctx.Response
	resp.Header.SetContentTypeBytes(contentTypeJson)
	resp.Header.SetStatusCode(fasthttp.StatusOK)
	p := string(ctx.Path())
	switch p {
	case "/status":
		resp.SetBody([]byte(`{ "status": "ok" }`))
		return
	case "/work":
		resp.SetBody(toJson(uuidV4(), fibonacciAt(20)))
		return
	}
	ctx.Error("404 page not found", fasthttp.StatusNotFound) // should Reset()
}

func main() {
	listenAddr := ":9090"
	server := &fasthttp.Server{
		Handler:                       handlerFunc,
		ReadBufferSize:                1024,
		WriteBufferSize:               512,
		ReadTimeout:                   4 * time.Second,
		WriteTimeout:                  4 * time.Second,
		IdleTimeout:                   8 * time.Second,
		GetOnly:                       true,
		NoDefaultServerHeader:         true,
		NoDefaultContentType:          true,
		NoDefaultDate:                 true,
		DisableHeaderNamesNormalizing: true,
	}
	_ = server.ListenAndServe(listenAddr)
}
