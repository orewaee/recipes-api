package controllers

import (
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
)

func (controller *RestController) getGuideById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	guide, err := controller.guideApi.GetGuideById(ctx, id)
	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	utils.MustWriteMarkdown(ctx, guide, fasthttp.StatusOK)
}
