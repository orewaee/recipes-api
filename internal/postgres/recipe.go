package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
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
		return nil, domain.ErrRecipeNotFound
	}

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (repo *RecipeRepo) GetRandomRecipe(ctx context.Context) (*domain.Recipe, error) {
	row := repo.pool.QueryRow(ctx, "select * from recipes order by random() limit 1")

	recipe := new(domain.Recipe)
	err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Description)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoRecipes
	}

	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func (repo *RecipeRepo) GetNumberOfRecipes(ctx context.Context) (int, error) {
	row := repo.pool.QueryRow(ctx, "select count(*) from recipes")

	number := 0
	if err := row.Scan(&number); err != nil {
		return 0, err
	}

	return number, nil
}

func (repo *RecipeRepo) GetRecipes(ctx context.Context, limit, offset int) ([]*domain.Recipe, error) {
	rows, err := repo.pool.Query(ctx, "select * from recipes limit $1 offset $2", limit, offset)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoRecipes
	}

	if err != nil {
		return nil, err
	}

	recipes, err := pgx.CollectRows[*domain.Recipe](rows, func(row pgx.CollectableRow) (*domain.Recipe, error) {
		recipe := new(domain.Recipe)
		if err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Description); err != nil {
			return nil, err
		}

		return recipe, nil
	})

	if err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return nil, domain.ErrNoRecipes
	}

	return recipes, nil
}

func (repo *RecipeRepo) GetRecipesByName(ctx context.Context, substring string, position domain.Position, limit, offset int) ([]*domain.Recipe, error) {
	sql := ""

	switch position {
	case domain.PositionStart:
		sql = fmt.Sprintf("select * from recipes where name ilike '%s%%' limit $1 offset $2", substring)
	case domain.PositionMiddle:
		sql = fmt.Sprintf("select * from recipes where name ilike '%%%s%%' limit $1 offset $2", substring)
	}

	rows, err := repo.pool.Query(ctx, sql, limit, offset)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoRecipes
	}

	if err != nil {
		return nil, err
	}

	recipes, err := pgx.CollectRows[*domain.Recipe](rows, func(row pgx.CollectableRow) (*domain.Recipe, error) {
		recipe := new(domain.Recipe)
		if err := row.Scan(&recipe.Id, &recipe.Name, &recipe.Description); err != nil {
			return nil, err
		}

		return recipe, nil
	})

	if err != nil {
		return nil, err
	}

	if len(recipes) == 0 {
		return nil, domain.ErrNoRecipes
	}

	return recipes, nil
}

func (repo *RecipeRepo) GetNameSuggestions(ctx context.Context, substring string, position domain.Position, limit int) ([]string, error) {
	sql := ""

	switch position {
	case domain.PositionStart:
		sql = fmt.Sprintf("select name from recipes where name ilike '%s%%' limit $1", substring)
	case domain.PositionMiddle:
		sql = fmt.Sprintf("select name from recipes where name ilike '%%%s%%' limit $1", substring)
	}

	rows, err := repo.pool.Query(ctx, sql, limit)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNoSuggestions
	}

	if err != nil {
		return nil, err
	}

	suggestions, err := pgx.CollectRows[string](rows, func(row pgx.CollectableRow) (string, error) {
		suggestion := ""
		if err := row.Scan(&suggestion); err != nil {
			return "", err
		}

		return suggestion, nil
	})

	if err != nil {
		return nil, err
	}

	if len(suggestions) == 0 {
		return nil, domain.ErrNoSuggestions
	}

	return suggestions, nil
}
