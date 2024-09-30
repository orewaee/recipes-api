package domain

import "errors"

var (
	ErrRecipeNotFound = errors.New("recipe not found")
	ErrNoRecipes      = errors.New("no recipes")
	ErrNoSuggestions  = errors.New("no suggestions")
	ErrNoKey          = errors.New("no key")
	ErrNoGuide        = errors.New("no guide")
	ErrNoPreview      = errors.New("no preview")
)
