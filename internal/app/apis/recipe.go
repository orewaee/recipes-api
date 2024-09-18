package apis

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/domain"
)

type RecipeApi interface {
	AddRecipe(ctx context.Context, name, description string) (*domain.Recipe, error)
	GetRecipeById(ctx context.Context, id string) (*domain.Recipe, error)
	GetRandomRecipe(ctx context.Context) (*domain.Recipe, error)
}
