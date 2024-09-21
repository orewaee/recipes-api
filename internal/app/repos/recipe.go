package repos

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/domain"
)

type RecipeRepo interface {
	AddRecipe(ctx context.Context, recipe *domain.Recipe) error
	GetRecipeById(ctx context.Context, id string) (*domain.Recipe, error)
	GetRandomRecipe(ctx context.Context) (*domain.Recipe, error)
	GetNumberOfRecipes(ctx context.Context) (int, error)
	GetRecipes(ctx context.Context, limit, offset int) ([]*domain.Recipe, error)
	GetNameSuggestions(ctx context.Context, substring string, position domain.Position, limit int) ([]string, error)
}
