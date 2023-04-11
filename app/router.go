package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nuraziz04/belajar-golang-restful-api/controller"
	"github.com/nuraziz04/belajar-golang-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler
	router.NotFound = http.HandlerFunc(exception.PageNotFoundHandler)

	return router
}
