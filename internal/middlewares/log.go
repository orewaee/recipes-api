package middlewares

import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

func LogMiddleware(handler fasthttp.RequestHandler, logger *zerolog.Logger) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		logger.Info().Str("route", string(ctx.RequestURI())).Send()
		handler(ctx)
	}
}
