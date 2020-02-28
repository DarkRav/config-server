package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
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
		log.Print(err.Error())
		writeErrorResponse(writer, 500, err.Error())
		return
	}
	writeResponse(writer, 200, configs)
}

func getRequestParamValue(request *http.Request, param string) string {
	vars := mux.Vars(request)
	value := vars[param]
	return value
}

func writeErrorResponse(writer http.ResponseWriter, status int, body string) {
	errorBody := make(map[string]string)
	errorBody["error"] = body
	errorBodyJson, _ := json.Marshal(errorBody)
	writeResponse(writer, 500, string(errorBodyJson))
}

func writeResponse(writer http.ResponseWriter, status int, body string) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	io.WriteString(writer, body)
}
