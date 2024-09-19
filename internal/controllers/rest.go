package controllers

import (
	fastrouter "github.com/fasthttp/router"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/valyala/fasthttp"
	"log"
)

type RestController struct {
	addr string
	api  apis.RecipeApi
}

func NewRestController(addr string, api apis.RecipeApi) *RestController {
	return &RestController{addr, api}
}

func (controller *RestController) Run() {
	router := fastrouter.New()

	router.GET("/recipe/{id}", controller.getRecipeById)
	router.GET("/recipe/random", controller.getRandomRecipe)
	router.GET("/recipes/number", controller.getNumberOfRecipes)

	if err := fasthttp.ListenAndServe(controller.addr, router.Handler); err != nil {
		log.Fatalln(err)
	}
}
