//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"github.com/nuraziz04/belajar-golang-restful-api/app"
	"github.com/nuraziz04/belajar-golang-restful-api/controller"
	"github.com/nuraziz04/belajar-golang-restful-api/middleware"
	"github.com/nuraziz04/belajar-golang-restful-api/repository"
	"github.com/nuraziz04/belajar-golang-restful-api/service"
)

// var categorySet = wire.NewSet(
// 	repository.NewCategoryRepository,
// 	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
// 	service.NewCategoryService,
// 	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
// 	controller.NewCategoryController,
// 	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
// )

func InitializedServer() *http.Server {

	wire.Build(
		app.NewDB,
		validator.New,
		repository.NewCategoryRepository,
		service.NewCategoryService,
		controller.NewCategoryController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddlware,
		NewServer,
	)
	return nil
}
