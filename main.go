package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"

	"github.com/nuraziz04/belajar-golang-restful-api/app"
	"github.com/nuraziz04/belajar-golang-restful-api/controller"
	"github.com/nuraziz04/belajar-golang-restful-api/helper"
	"github.com/nuraziz04/belajar-golang-restful-api/middleware"
	"github.com/nuraziz04/belajar-golang-restful-api/repository"
	"github.com/nuraziz04/belajar-golang-restful-api/service"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	authMiddleware := middleware.NewAuthMiddlware(router)

	server := NewServer(authMiddleware)

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
