package middlewares

import "github.com/valyala/fasthttp"

var (
	corsAllowHeaders     = "Authorization, Content-Type, Accept"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS,PATCH"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

func CorsMiddleware(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)

		if string(ctx.Method()) == fasthttp.MethodOptions {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}

		handler(ctx)
	}
}
