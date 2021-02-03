package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/valyala/fasthttp"
)

var started uint64 = 0
var completed uint64 = 0
var delay time.Duration

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("args")
	}
	address := os.Args[1]
	ms, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalln("delay", err)
	}

	delay = time.Millisecond * time.Duration(ms)

	h := requestHandler
	if err := fasthttp.ListenAndServe(address, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	atomic.AddUint64(&started, 1)

	delta := started - completed
	fmt.Fprintf(ctx, "%d - %d = %d\n\n", started, completed, delta)
	if delay > 0 {
		time.Sleep(delay)
	}
	atomic.AddUint64(&completed, 1)
}
