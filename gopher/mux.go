package main

import "github.com/valyala/fasthttp"

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

type mux struct {
	callbackC chan<- string // write-only
}

func (m *mux) handlerFunc(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.SetContentTypeBytes(contentTypeJson)
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	p := string(ctx.Path())
	switch p {
	case pathStatus:
		ctx.Response.SetBody(responseOK)
		return
	case pathTest:
		ctx.Response.SetBody(responseOK) //todo specs mofo
		return
	case pathCombat:
		ctx.Response.SetBody(responseOK)
		m.callbackC <- pathJab
		m.callbackC <- pathHook
		return
	case pathJab:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(2)))
		m.callbackC <- pathJab
		m.callbackC <- pathJab
		return
	case pathCross:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(4)))
		m.callbackC <- pathJab
		m.callbackC <- pathJab
		m.callbackC <- pathCross
		return
	case pathHook:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(8)))
		m.callbackC <- pathHook
		m.callbackC <- pathHook
		m.callbackC <- pathUppercut
		return
	case pathUppercut:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(16)))
		m.callbackC <- pathCross
		m.callbackC <- pathHook
		m.callbackC <- pathUppercut
		return
	case pathWork:
		ctx.Response.SetBody(toJson(uuidV4(), fibonacciAt(32 /* really unhappy FormatUint */)))
		return
	}
	ctx.Error("404 page not found", fasthttp.StatusNotFound) // should Reset()
}
