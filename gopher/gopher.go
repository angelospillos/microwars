package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

const (
	pathStatus   = "/status"
	pathTest     = "/test"
	pathCombat   = "/combat"
	pathJab      = "/jab"
	pathCross    = "/cross"
	pathHook     = "/hook"
	pathUppercut = "/uppercut"
	pathWork     = "/work"
)

var (
	contentTypeJson = []byte("application/json")
	responseOK      = []byte(`{"status":"ok"}`)
)

func handlerFunc(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentTypeBytes(contentTypeJson)
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	p := string(ctx.Path())
	switch p {
	case pathStatus:
		ctx.Response.SetBody(responseOK)
		return
	case pathTest:
		ctx.Response.SetBody(responseOK)
		return
	case pathCombat:
		ctx.Response.SetBody(responseOK)
		return
	case pathJab:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(2)))
		return
	case pathCross:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(4)))
		return
	case pathHook:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(8)))
		return
	case pathUppercut:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(16)))
		return
	case pathWork:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(32 /* really unhappy FormatUint */)))
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
	log.Printf("Listening on %s", listenAddr)
	log.Fatal(server.ListenAndServe(listenAddr))
}
