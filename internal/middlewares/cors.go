package middlewares

import "github.com/valyala/fasthttp"

func CorsMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
		ctx.Response.Header.SetBytesV("Access-Control-Allow-Origin", ctx.Request.Header.Peek("Origin"))
		handler(ctx)
	}
}
