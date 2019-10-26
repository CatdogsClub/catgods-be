// @title CatDogs API
// @version 1.0
// @description CatDogs API 文档
// @contact.name Yoko
// @BasePath /api

package main

import (
	_ "catdogs-be/docs"
	"catdogs-be/models"
	"catdogs-be/routers"
	"fmt"
	"net/http"
	"time"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	models.InitModel()
}

func main() {
	runServer()
}

func runServer() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           ":9999",
		Handler:        router,
		ReadTimeout:    18 * time.Second,
		WriteTimeout:   18 * time.Second,
		MaxHeaderBytes: 1 << 28,
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
