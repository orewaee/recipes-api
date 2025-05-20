package controllers

import (
	"log"

	fastrouter "github.com/fasthttp/router"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/middlewares"
	"github.com/orewaee/recipes-api/internal/utils"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
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

	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		utils.MustWriteString(ctx, "pong", fasthttp.StatusOK)
	})

	router.GET("/recipe/{id}", controller.getRecipeById)
	router.GET("/recipe/random", controller.getRandomRecipe)
	router.POST("/recipe", controller.postRecipe)

	router.GET("/recipes/number", controller.getNumberOfRecipes)
	router.GET("/recipes", controller.getRecipes)
	router.GET("/recipes/suggestions", controller.getNameSuggestions)

	router.GET("/guide/{id}", controller.getGuideById)
	router.POST("/guide/{id}", controller.postGuide)

	router.GET("/preview/{id}", controller.getPreviewById)
	router.POST("/preview/{id}", controller.postPreview)

	controller.logger.Info().Msgf("running on addr %s", controller.addr)

	if err := fasthttp.ListenAndServe(controller.addr, middlewares.LogMiddleware(middlewares.CorsMiddleware(router.Handler), controller.logger)); err != nil {
		log.Fatalln(err)
	}
}
