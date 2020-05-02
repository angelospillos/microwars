package main

import (
	"github.com/valyala/fasthttp"
	"time"
)

const requestTimeout = 4 * time.Second

func doDeadlineGET(c *fasthttp.Client, requestURI string) error {
	req := fasthttp.AcquireRequest() // recycle and reduce GC pressure
	req.SetRequestURI(requestURI)
	req.Header.SetMethod(fasthttp.MethodGet)
	req.Header.SetBytesV("Accept", contentTypeJson)
	err := c.DoDeadline(req, nil /* ignore response */, time.Now().Add(requestTimeout))
	fasthttp.ReleaseRequest(req)
	return err
}
