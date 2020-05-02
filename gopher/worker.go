package main

import (
	"log"
	"sync"
)

var workerPoolInit sync.Once

type workerPool struct {
	workRequestC   <-chan string // read-only
	workHandleFunc func(string) error
}

func (p *workerPool) poolInit(w int) {
	workerPoolInit.Do(func() {
		for i := 0; i < w; i++ {
			go p.workerInit()
		}
	})
}

func (p *workerPool) workerInit() {
	var err error
	for w := range p.workRequestC {
		err = p.workHandleFunc(w)
		if err != nil {
			log.Println(err) //todo error channel
		}
	}
}
