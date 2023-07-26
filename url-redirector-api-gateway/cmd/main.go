package main

import (
	"log"

	"url-redirector-api-gateway/docs"
	"url-redirector-api-gateway/pkg/auth"
	"url-redirector-api-gateway/pkg/config"
	"url-redirector-api-gateway/pkg/url"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to config", err)
	}

	routes := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	routes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	authSvc := *auth.RegisterRoutes(routes, &c)
	url.RegisterRoutes(routes, &c, &authSvc)

	routes.Run(c.Port)
}
