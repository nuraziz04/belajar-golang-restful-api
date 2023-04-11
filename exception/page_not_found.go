package exception

import (
	"net/http"

	"github.com/nuraziz04/belajar-golang-restful-api/helper"
	"github.com/nuraziz04/belajar-golang-restful-api/model/web"
)

func PageNotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	pageNotFound(writer, request)
}

func pageNotFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code:   http.StatusNotFound,
		Status: "PAGE NOT FOUND",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
