package services

import (
	"context"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
)

type RecipeService struct {
	repo repos.RecipeRepo
}

func NewRecipeService(repo repos.RecipeRepo) apis.RecipeApi {
	return &RecipeService{repo}
}

func (service *RecipeService) AddRecipe(ctx context.Context, name, description string) (*domain.Recipe, error) {
	id := gonanoid.Must(8)

	recipe := &domain.Recipe{
		Id:          id,
		Name:        name,
		Description: description,
	}

	if err := service.repo.AddRecipe(ctx, recipe); err != nil {
		return nil, err
	}

	return recipe, nil
}

func (service *RecipeService) GetRecipeById(ctx context.Context, id string) (*domain.Recipe, error) {
	return service.repo.GetRecipeById(ctx, id)
}

func (service *RecipeService) GetRandomRecipe(ctx context.Context) (*domain.Recipe, error) {
	return service.repo.GetRandomRecipe(ctx)
}

func (service *RecipeService) GetNumberOfRecipes(ctx context.Context) (int, error) {
	return service.repo.GetNumberOfRecipes(ctx)
}
