package main

import (
	"net/http"
)
import "github.com/gorilla/mux"

func main() {
	environmentService := createEnvironmentService()
	configController := &ConfigController{
		environmentService: *environmentService,
	}
	router := mux.NewRouter()
	router.HandleFunc("/{application}-{profile}.yml", configController.GetConfigs).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
