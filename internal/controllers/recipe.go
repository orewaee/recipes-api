package controllers

import (
	"errors"
	"github.com/orewaee/recipes-api/internal/constants"
	"github.com/orewaee/recipes-api/internal/dtos"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
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
