package controllers

import (
	"errors"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
)

func (controller *RestController) postPreview(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	if err := controller.previewApi.AddPreview(ctx, id, ctx.PostBody()); err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	utils.MustWriteString(ctx, id, fasthttp.StatusOK)
}

func (controller *RestController) getPreviewById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	preview, err := controller.previewApi.GetPreviewById(ctx, id)
	if err != nil && errors.Is(err, domain.ErrNoPreview) {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusNotFound)
		return
	}

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set(fasthttp.HeaderContentType, "image/png")
	utils.MustWriteBytes(ctx, preview, fasthttp.StatusOK)
}
