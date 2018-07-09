package main

import (
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("Hello"))
}

func main() {
	fasthttp.ListenAndServe(":5003", fastHTTPHandler)
}
