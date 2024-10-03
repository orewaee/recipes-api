package domain

import "errors"

var (
	ErrRecipeNotFound    = errors.New("recipe not found")
	ErrNoRecipes         = errors.New("no recipes")
	ErrNoSuggestions     = errors.New("no suggestions")
	ErrNoKey             = errors.New("no key")
	ErrNoGuide           = errors.New("no guide")
	ErrNoPreview         = errors.New("no preview")
	ErrRecipeName        = errors.New("the recipe name must be between 1 and 64 in length")
	ErrRecipeDescription = errors.New("the recipe description must be between 1 and 256 in length")
)
