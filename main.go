package main

import (
	"fmt"
	"net/http"

	. "auth-service/configs"
	"auth-service/logger"
	"auth-service/middlewares"
	. "auth-service/router"

	"github.com/gorilla/mux"
)

func main() {
	logger.InitLogger()
	config := GetAppConfig()
	muxRouter := mux.NewRouter()

	CreateRoutes(muxRouter)
	muxRouter.Use(middlewares.LoggingMiddleware)
	muxRouter.Use(middlewares.WriteApplicationJsonMiddleware)

	logger.Infof("Server is starting on http://localhost:%v...", config.PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), muxRouter)
	if err != nil {
		logger.Error(err)
	}
}
