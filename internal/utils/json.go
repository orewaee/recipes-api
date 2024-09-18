package utils

import (
	"github.com/valyala/fasthttp"
)

func MustWriteJson(ctx *fasthttp.RequestCtx, data interface{}, code int) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	MustWriteAny(ctx, data, code)
}
