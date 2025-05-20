package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/dtos"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/valyala/fasthttp"
)

func (controller *RestController) postRecipe(ctx *fasthttp.RequestCtx) {
	dtoRequest := new(dtos.RecipeRequest)
	if err := json.Unmarshal(ctx.PostBody(), dtoRequest); err != nil {
		utils.MustWriteString(ctx, "failed to parse the request body", fasthttp.StatusInternalServerError)
		return
	}

	if err := dtoRequest.Validate(); err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}

	recipe, err := controller.recipeApi.AddRecipe(ctx, dtoRequest.Name, dtoRequest.Description)
	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	dtoRecipe := dtos.Recipe{
		Id:          recipe.Id,
		Name:        recipe.Name,
		Description: recipe.Description,
	}

	utils.MustWriteJson(ctx, dtoRecipe, fasthttp.StatusCreated)
}

func (controller *RestController) getRecipeById(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	recipe, err := controller.recipeApi.GetRecipeById(ctx, id)
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
	recipe, err := controller.recipeApi.GetRandomRecipe(ctx)

	if err != nil && errors.Is(err, domain.ErrNoRecipes) {
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
	number, err := controller.recipeApi.GetNumberOfRecipes(ctx)
	fmt.Println(number)

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

	perPage, err := strconv.Atoi(string(ctx.QueryArgs().Peek("per_page")))
	if err != nil {
		utils.MustWriteString(ctx, "invalid per_page", fasthttp.StatusBadRequest)
		return
	}

	name := string(ctx.QueryArgs().Peek("name"))
	if name != "" {
		recipes, err := controller.recipeApi.GetRecipesByName(ctx, name, domain.PositionStart, perPage, page)
		if err != nil && !errors.Is(err, domain.ErrNoSuggestions) {
			utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
			return
		}

		if err == nil {
			utils.MustWriteJson(ctx, recipes, fasthttp.StatusOK)
			return
		}

		recipes, err = controller.recipeApi.GetRecipesByName(ctx, name, domain.PositionMiddle, perPage, page)

		if err != nil && errors.Is(err, domain.ErrNoSuggestions) {
			utils.MustWriteString(ctx, err.Error(), fasthttp.StatusNotFound)
			return
		}

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
		return
	}

	recipes, err := controller.recipeApi.GetRecipes(ctx, perPage, page)
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

func (controller *RestController) getNameSuggestions(ctx *fasthttp.RequestCtx) {
	query := string(ctx.QueryArgs().Peek("query"))

	limit, err := strconv.Atoi(string(ctx.QueryArgs().Peek("limit")))
	if err != nil {
		utils.MustWriteString(ctx, "invalid limit", fasthttp.StatusBadRequest)
		return
	}

	suggestions, err := controller.recipeApi.GetNameSuggestions(ctx, query, domain.PositionStart, limit)
	if err != nil && !errors.Is(err, domain.ErrNoSuggestions) {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	if err == nil {
		response := make([]dtos.Suggestion, len(suggestions))
		for i := range suggestions {
			response[i] = dtos.Suggestion{
				Id:   suggestions[i].Id,
				Name: suggestions[i].Name,
			}
		}

		utils.MustWriteJson(ctx, response, fasthttp.StatusOK)
		return
	}

	suggestions, err = controller.recipeApi.GetNameSuggestions(ctx, query, domain.PositionMiddle, limit)

	if err != nil && errors.Is(err, domain.ErrNoSuggestions) {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusNotFound)
		return
	}

	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	response := make([]dtos.Suggestion, len(suggestions))
	for i := range suggestions {
		response[i] = dtos.Suggestion{
			Id:   suggestions[i].Id,
			Name: suggestions[i].Name,
		}
	}

	utils.MustWriteJson(ctx, response, fasthttp.StatusOK)
}

func (controller *RestController) updateRecipe(ctx *fasthttp.RequestCtx) {
	id := ctx.UserValue("id").(string)
	if id == "" {
		utils.MustWriteString(ctx, "missing id", fasthttp.StatusOK)
		return
	}

	data := utils.MustReadJson[dtos.UpdateRecipe](ctx)
	var name, description *string

	if data.Name != "" {
		name = &data.Name
	}

	if data.Description != "" {
		description = &data.Description
	}

	err := controller.recipeApi.UpdateRecipe(ctx, id, name, description)
	if err != nil {
		utils.MustWriteString(ctx, err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	utils.MustWriteString(ctx, id, fasthttp.StatusOK)
}
