package controllers

import (
	fastrouter "github.com/fasthttp/router"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/middlewares"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"log"
)

type RestController struct {
	addr       string
	recipeApi  apis.RecipeApi
	guideApi   apis.GuideApi
	previewApi apis.PreviewApi
	logger     *zerolog.Logger
}

func NewRestController(addr string, recipeApi apis.RecipeApi, guideApi apis.GuideApi, previewApi apis.PreviewApi, logger *zerolog.Logger) *RestController {
	return &RestController{addr, recipeApi, guideApi, previewApi, logger}
}

func (controller *RestController) Run() {
	router := fastrouter.New()

	router.GET("/recipe/{id}", middlewares.LogMiddleware(controller.getRecipeById, controller.logger))
	router.GET("/recipe/random", middlewares.LogMiddleware(controller.getRandomRecipe, controller.logger))
	router.POST("/recipe", middlewares.LogMiddleware(controller.postRecipe, controller.logger))

	router.GET("/recipes/number", middlewares.LogMiddleware(controller.getNumberOfRecipes, controller.logger))
	router.GET("/recipes", middlewares.LogMiddleware(controller.getRecipes, controller.logger))
	router.GET("/recipes/suggestions", middlewares.LogMiddleware(controller.getNameSuggestions, controller.logger))

	router.GET("/guide/{id}", middlewares.LogMiddleware(controller.getGuideById, controller.logger))
	router.POST("/guide/{id}", middlewares.LogMiddleware(controller.postGuide, controller.logger))

	router.GET("/preview/{id}", middlewares.LogMiddleware(controller.getPreviewById, controller.logger))
	router.POST("/preview/{id}", middlewares.LogMiddleware(controller.postPreview, controller.logger))

	if err := fasthttp.ListenAndServe(controller.addr, router.Handler); err != nil {
		log.Fatalln(err)
	}
}
