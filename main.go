package main

import (
	"encoding/json"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/tVienonen/recipe-manager-backend/middleware/config_provider"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	app := iris.New()
	debug := os.Getenv("DEBUG")
	if debug == "TRUE" {
		app.Logger().SetLevel("debug")
	} else {
		app.Logger().SetLevel("info")
	}
	config := &config_provider.DBConfig{}
	configData := os.Getenv("DB_CONFIG")
	json.Unmarshal([]byte(configData), &config)
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(func(ctx iris.Context) {
		ctx.Values().Set("validate", validate)
		ctx.Next()
	})
	app.Use(config_provider.New(config))
	RegisterRoutes(app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
