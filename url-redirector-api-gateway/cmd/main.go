package main

import (
	"log"
	"url-redirector-api-gateway/pkg/auth"
	"url-redirector-api-gateway/pkg/config"
	"url-redirector-api-gateway/pkg/url"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed to config", err)
	}

	routes := gin.Default()

	authSvc := *auth.RegisterRoutes(routes, &c)
	url.RegisterRoutes(routes, &c, &authSvc)

	routes.Run(c.Port)
}
