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

	requestTimeout = 4 * time.Second

	workerCount  = 256 // approx. 1MB pre-allocated stack space
	workerBuffer = 8
)

var (
	contentTypeJson = []byte("application/json")
	responseOK      = []byte(`{"status":"ok"}`)
)

func handlerFunc(callbacks chan string) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.SetContentTypeBytes(contentTypeJson)
		ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
		p := string(ctx.Path())
		switch p {
		case pathStatus:
			ctx.Response.SetBody(responseOK)
			return
		case pathTest:
			ctx.Response.SetBody(responseOK)
			callbacks <- pathStatus
			return
		case pathCombat:
			ctx.Response.SetBody(responseOK)
			callbacks <- pathJab
			callbacks <- pathHook
			return
		case pathJab:
			ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(2)))
			callbacks <- pathJab
			callbacks <- pathJab
			return
		case pathCross:
			ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(4)))
			callbacks <- pathJab
			callbacks <- pathJab
			callbacks <- pathCross
			return
		case pathHook:
			ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(8)))
			callbacks <- pathHook
			callbacks <- pathHook
			callbacks <- pathUppercut
			return
		case pathUppercut:
			ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(16)))
			callbacks <- pathCross
			callbacks <- pathHook
			callbacks <- pathUppercut
			return
		case pathWork:
			ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(32 /* really unhappy FormatUint */)))
			return
		}
		ctx.Error("404 page not found", fasthttp.StatusNotFound) // should Reset()
	}
}

func main() {
	listenAddr := getEnv("LISTEN_ADDR", ":9090")
	targetAddr := requireEnv("TARGET_ADDR")

	callbacks := make(chan string, workerBuffer /* buffered: avoid thrashing */)
	server := &fasthttp.Server{
		Handler:                       handlerFunc(callbacks),
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
	for i := 0; i < workerCount; i++ { // worker pool
		go func() {
			for requestPath := range callbacks {
				req := fasthttp.AcquireRequest() // recycle and reduce GC pressure
				req.Header.SetMethod(fasthttp.MethodGet)
				req.Header.SetHost(targetAddr)
				req.SetRequestURI(requestPath)
				err := client.DoDeadline(req, nil /* ignore response */, time.Now().Add(requestTimeout))
				if err != nil { //todo ErrTimeout vs ?
					log.Panicf("Error for HTTP '%s' request: %v", requestPath, err)
				}
				fasthttp.ReleaseRequest(req)
			}
		}()
	}
	log.Printf("Listening on %s", listenAddr)
	log.Fatal(server.ListenAndServe(listenAddr))
}
