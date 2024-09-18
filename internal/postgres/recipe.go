package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"github.com/orewaee/recipes-api/internal/constants"
)

type RecipeRepo struct {
	pool *pgxpool.Pool
}

func NewRecipeRepo(ctx context.Context, connString string) (repos.RecipeRepo, error) {
	pool, err := NewPool(ctx, connString)
	if err != nil {
		return nil, err
	}

	sql := `
	create table if not exists recipes (
	    id char(8) primary key,
	    name varchar(64) not null,
	    description varchar(256)
	)
	`

	if _, err := pool.Exec(ctx, sql); err != nil {
		return nil, err
	}

	return &RecipeRepo{pool}, nil
}

func (repo *RecipeRepo) AddRecipe(ctx context.Context, recipe *domain.Recipe) error {
	sql := "insert into recipes (id, name, description) values ($1, $2, $3)"
	_, err := repo.pool.Exec(ctx, sql, recipe.Id, recipe.Name, recipe.Description)

	return err
}

func (repo *RecipeRepo) GetRecipeById(ctx context.Context, id string) (*domain.Recipe, error) {
	row := repo.pool.QueryRow(ctx, "select * from recipes where id = $1", id)

	recipe := new(domain.Recipe)
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Description)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, constants.ErrRecipeNotFound
	}

	if err != nil {
		return nil, err
	}

	return recipe, nil
}
