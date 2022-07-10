package main

import (
	"DemoSite/controller"
	_ "DemoSite/docs"
	"DemoSite/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
func main() {
	server := gin.Default()
	models.CreateDataBase()
	controller.UserControllerInit(server)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	server.Run()
}
