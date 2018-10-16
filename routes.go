package main

import (
	"github.com/kataras/iris"
	"github.com/tVienonen/recipe-manager-backend/handlers/recipes"
)

func RegisterRoutes(app *iris.Application) {
	apiRoutes := app.Party("/api")
	apiV1Routes := apiRoutes.Party("/v1")
	apiV1Routes.Get("/recipes", recipes.GetRecipesHandler)
}
