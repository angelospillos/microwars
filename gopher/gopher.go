package main

import (
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

const workerCount = 256 //todo as flag

var pipelineC = make(chan string /* no buffering: thrashing over blocking */)

func main() {
	listenAddr := getEnv("LISTEN_ADDR", ":9090")
	targetAddr := requireEnv("TARGET_ADDR")
	router := &mux{callbackC: pipelineC}
	server := &fasthttp.Server{
		Handler:                       router.handlerFunc,
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
	client := &fasthttp.Client{
		NoDefaultUserAgentHeader:      true,
		MaxConnsPerHost:               1024,
		MaxIdemponentCallAttempts:     2,
		MaxResponseBodySize:           1024,
		ReadBufferSize:                1024,
		WriteBufferSize:               512,
		ReadTimeout:                   4 * time.Second,
		WriteTimeout:                  4 * time.Second,
		MaxConnWaitTimeout:            8 * time.Second,
		MaxIdleConnDuration:           8 * time.Second,
		DisableHeaderNamesNormalizing: true,
		DisablePathNormalizing:        true,
	}
	workers := &workerPool{
		workRequestC: pipelineC,
		workHandleFunc: func(requestPath string) error {
			return doDeadlineGET(client, targetAddr+requestPath)
		},
	}

	log.Printf("Using %d background workers", workerCount)
	workers.poolInit(workerCount)

	log.Printf("Listening on %s", listenAddr)
	log.Fatal(server.ListenAndServe(listenAddr))
}
