package controllers

import (
	"errors"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
)

func (controller *RestController) postGuide(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	markdown := string(ctx.PostBody())
	if err := controller.guideApi.AddGuide(ctx, id, markdown); err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	utils.MustWriteMarkdown(ctx, id, fasthttp.StatusOK)
}

func (controller *RestController) getGuideById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	guide, err := controller.guideApi.GetGuideById(ctx, id)
	if err != nil && errors.Is(err, domain.ErrNoGuide) {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusNotFound)
		return
	}

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.Response.Header.Set("Cache-Control", "max-age=600")
	utils.MustWriteMarkdown(ctx, guide, fasthttp.StatusOK)
}
