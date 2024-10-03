package dtos

import "github.com/orewaee/recipes-api/internal/app/domain"

type Recipe struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RecipeRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (request *RecipeRequest) Validate() error {
	nameLength := len(request.Name)
	if nameLength < 0 || nameLength > 64 {
		return domain.ErrRecipeName
	}

	descriptionLength := len(request.Description)
	if descriptionLength < 1 || descriptionLength > 256 {
		return domain.ErrRecipeDescription
	}

	return nil
}
