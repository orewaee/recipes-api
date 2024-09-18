package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func MustWriteBytes(ctx *fasthttp.RequestCtx, data []byte, code int) {
	ctx.SetStatusCode(code)
	if _, err := ctx.Write(data); err != nil {
		panic(err)
	}
}

func MustWriteString(ctx *fasthttp.RequestCtx, data string, code int) {
	ctx.SetStatusCode(code)
	if _, err := ctx.WriteString(data); err != nil {
		panic(err)
	}
}

func MustWriteAny(ctx *fasthttp.RequestCtx, data any, code int) {
	ctx.SetStatusCode(code)

	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if _, err := ctx.Write(bytes); err != nil {
		panic(err)
	}
}
