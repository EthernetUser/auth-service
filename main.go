package main

import (
	. "auth-service/configs"
	"auth-service/logger"
	"auth-service/middlewares"
	. "auth-service/router"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
