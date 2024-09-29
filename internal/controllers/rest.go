package controllers

import (
	fastrouter "github.com/fasthttp/router"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/valyala/fasthttp"
	"log"
)

type RestController struct {
	addr       string
	recipeApi  apis.RecipeApi
	guideApi   apis.GuideApi
	previewApi apis.PreviewApi
}

func NewRestController(addr string, recipeApi apis.RecipeApi, guideApi apis.GuideApi, previewApi apis.PreviewApi) *RestController {
	return &RestController{addr, recipeApi, guideApi, previewApi}
}

func (controller *RestController) Run() {
	router := fastrouter.New()

	router.GET("/recipe/{id}", controller.getRecipeById)
	router.GET("/recipe/random", controller.getRandomRecipe)

	router.GET("/recipes/number", controller.getNumberOfRecipes)
	router.GET("/recipes", controller.getRecipes)
	router.GET("/recipes/suggestions", controller.getNameSuggestions)

	router.GET("/guide/{id}", controller.getGuideById)

	router.GET("/preview/{id}", controller.getPreviewById)
	router.POST("/preview/{id}", controller.postPreview)

	if err := fasthttp.ListenAndServe(controller.addr, router.Handler); err != nil {
		log.Fatalln(err)
	}
}
