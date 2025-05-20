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
	GetRecipesByName(ctx context.Context, substring string, position domain.Position, limit, offset int) ([]*domain.Recipe, error)
	GetNameSuggestions(ctx context.Context, substring string, position domain.Position, limit int) ([]domain.Suggestion, error)
	SetRecipeName(ctx context.Context, id, name string) error
	SetRecipeDescription(ctx context.Context, id, description string) error
}
