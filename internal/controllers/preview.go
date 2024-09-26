package controllers

import (
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
)

func (controller *RestController) getPreviewById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	preview, err := controller.previewApi.GetPreviewById(ctx, id)

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set(fasthttp.HeaderContentType, "image/png")
	utils.MustWriteBytes(ctx, preview, fasthttp.StatusOK)
}
