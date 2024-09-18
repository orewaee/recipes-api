package constants

import "errors"

var (
	ErrRecipeNotFound = errors.New("recipe not found")
	ErrNoRecipes      = errors.New("no recipes")
)
