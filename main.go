package main

import (
	"DemoSite/controller"
	_ "DemoSite/docs"
	jwtDomain "DemoSite/domain/jwt"
	"DemoSite/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin Swagger Demo
// @version 1.0
// @description Swagger API.
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	server := gin.Default()
	models.CreateDataBase()
	setJwtSetting()
	controller.UserControllerInit(server)
	controller.AuthControllerInit(server)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	server.Run()
}

func setJwtSetting() {
	setting := jwtDomain.JwtSetting{
		"a1234567890",
		"b1234567890",
		30,
	}
	jwtDomain.SetJwtSetting(&setting)
}
