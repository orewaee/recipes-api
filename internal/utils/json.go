package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func MustWriteJson(ctx *fasthttp.RequestCtx, data interface{}, code int) {
	ctx.Response.Header.Set("Content-Type", "application/json")
	MustWriteAny(ctx, data, code)
}

func MustReadJson[T interface{}](ctx *fasthttp.RequestCtx) *T {
	data := new(T)
	if err := json.Unmarshal(ctx.Request.Body(), data); err != nil {
		return nil
	}

	return data
}
