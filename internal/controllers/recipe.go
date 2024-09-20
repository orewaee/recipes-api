package controllers

import (
	"errors"
	"github.com/orewaee/recipes-api/internal/constants"
	"github.com/orewaee/recipes-api/internal/dtos"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
	"strconv"
)

func (controller *RestController) getRecipeById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)

	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	recipe, err := controller.api.GetRecipeById(ctx, id)
	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusOK)
		return
	}

	recipeDto := &dtos.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
	}

	utils.MustWriteJson(ctx, recipeDto, fasthttp.StatusOK)
}

func (controller *RestController) getRandomRecipe(ctx *fasthttp.RequestCtx) {
	recipe, err := controller.api.GetRandomRecipe(ctx)

	if err != nil && errors.Is(err, constants.ErrNoRecipes) {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusNotFound)
		return
	}

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	recipeDto := &dtos.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
	}

	utils.MustWriteJson(ctx, recipeDto, fasthttp.StatusOK)
}

func (controller *RestController) getNumberOfRecipes(ctx *fasthttp.RequestCtx) {
	number, err := controller.api.GetNumberOfRecipes(ctx)

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	utils.MustWriteAny(ctx, number, fasthttp.StatusOK)
}

func (controller *RestController) getRecipes(ctx *fasthttp.RequestCtx) {
	page, err := strconv.Atoi(string(ctx.QueryArgs().Peek("page")))
	if err != nil {
		utils.MustWriteString(ctx, "invalid page", fasthttp.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		utils.MustWriteString(ctx, "invalid limit", fasthttp.StatusBadRequest)
		return
	}

	recipes, err := controller.api.GetRecipes(ctx, limit, page)
	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	dtoRecipes := make([]*dtos.Recipe, len(recipes))
	for i, recipe := range recipes {
		dtoRecipes[i] = &dtos.Recipe{
			Id:          recipe.Id,
			Name:        recipe.Name,
			Description: recipe.Description,
		}
	}

	utils.MustWriteJson(ctx, dtoRecipes, fasthttp.StatusOK)
}
