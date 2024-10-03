package services

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/rs/zerolog"
	"strconv"
	"time"
)

type RecipeService struct {
	recipeRepo repos.RecipeRepo
	cacheRepo  repos.CacheRepo
	logger     *zerolog.Logger
}

func NewRecipeService(recipeRepo repos.RecipeRepo, cacheRepo repos.CacheRepo, logger *zerolog.Logger) apis.RecipeApi {
	return &RecipeService{recipeRepo, cacheRepo, logger}
}

func (service *RecipeService) AddRecipe(ctx context.Context, name, description string) (*domain.Recipe, error) {
	id := gonanoid.Must(8)

	recipe := &domain.Recipe{
		Id:          id,
		Name:        name,
		Description: description,
	}

	service.logger.Log().
		Str("id", id).
		Str("name", name).
		Msg("new recipe")

	if err := service.recipeRepo.AddRecipe(ctx, recipe); err != nil {
		service.logger.Error().Err(err).Send()
		return nil, err
	}

	return recipe, nil
}

func (service *RecipeService) GetRecipeById(ctx context.Context, id string) (*domain.Recipe, error) {
	key := "recipe_" + id

	if cache, err := service.cacheRepo.Get(ctx, key); err == nil {
		return utils.MustUnmarshalRecipe(cache), nil
	}

	recipe, err := service.recipeRepo.GetRecipeById(ctx, id)
	if err != nil {
		return nil, err
	}

	data := utils.MustMarshalRecipe(recipe)
	service.cacheRepo.Put(ctx, key, data, time.Minute*5)

	return recipe, nil
}

func (service *RecipeService) GetRandomRecipe(ctx context.Context) (*domain.Recipe, error) {
	return service.recipeRepo.GetRandomRecipe(ctx)
}

func (service *RecipeService) GetNumberOfRecipes(ctx context.Context) (int, error) {
	key := "recipes_number"

	if value, err := service.cacheRepo.Get(ctx, key); err == nil {
		service.logger.Error().Err(err).Send()
		return strconv.Atoi(value)
	}

	number, err := service.recipeRepo.GetNumberOfRecipes(ctx)
	if err != nil {
		service.logger.Error().Err(err).Send()
		return 0, err
	}

	service.cacheRepo.Put(ctx, key, strconv.Itoa(number), time.Minute*5)

	return service.recipeRepo.GetNumberOfRecipes(ctx)
}

func (service *RecipeService) GetRecipes(ctx context.Context, limit, page int) ([]*domain.Recipe, error) {
	offset := limit * (page - 1)

	recipes, err := service.recipeRepo.GetRecipes(ctx, limit, offset)
	if err != nil {
		service.logger.Error().Err(err).Send()
		return nil, err
	}

	return recipes, nil
}

func (service *RecipeService) GetRecipesByName(ctx context.Context, substring string, position domain.Position, limit, page int) ([]*domain.Recipe, error) {
	offset := limit * (page - 1)

	recipes, err := service.recipeRepo.GetRecipesByName(ctx, substring, position, limit, offset)
	if err != nil {
		service.logger.Error().Err(err).Send()
		return nil, err
	}

	return recipes, nil
}

func (service *RecipeService) GetNameSuggestions(ctx context.Context, substring string, position domain.Position, limit int) ([]string, error) {
	return service.recipeRepo.GetNameSuggestions(ctx, substring, position, limit)
}
