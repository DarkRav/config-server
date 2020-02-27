package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type ConfigController struct {
	environmentService EnvironmentService
}

func (controller *ConfigController) GetConfigs(writer http.ResponseWriter, request *http.Request) {
	application := getRequestParamValue(request, "application")
	profile := getRequestParamValue(request, "profile")

	configs, err := controller.environmentService.getConfigs(application, profile)
	if err != nil {
		writeResponse(writer, 500, err.Error())
		return
	}
	writeResponse(writer, 200, configs)
}

func getRequestParamValue(request *http.Request, param string) string {
	vars := mux.Vars(request)
	value := vars[param]
	return value
}

func writeResponse(writer http.ResponseWriter, status int, body string) {
	writer.WriteHeader(status)
	io.WriteString(writer, body)
}
