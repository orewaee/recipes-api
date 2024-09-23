package utils

import (
	"encoding/json"
	"github.com/orewaee/recipes-api/internal/app/domain"
)

func MustUnmarshalRecipe(data string) *domain.Recipe {
	recipe := new(domain.Recipe)
	if err := json.Unmarshal([]byte(data), recipe); err != nil {
		panic(err)
	}

	return recipe
}

func MustMarshalRecipe(recipe *domain.Recipe) string {
	bytes, err := json.Marshal(recipe)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
