package url

import (
	"url-redirector-api-gateway/pkg/auth"
	"url-redirector-api-gateway/pkg/config"
	"url-redirector-api-gateway/pkg/url/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	r.GET("/urls/:id", svc.GetUrl)
	r.Use(a.AuthRequired)
	r.GET("/urls", svc.GetUserUrls)
	r.POST("/urls", svc.AddUrl)
	r.POST("/urls/activate", svc.ActivateUrl)

	return svc
}

func (svc *ServiceClient) GetUrl(ctx *gin.Context) {
	routes.GetUrl(ctx, svc.Client)
}

func (svc *ServiceClient) AddUrl(ctx *gin.Context) {
	routes.AddUrl(ctx, svc.Client)
}

func (svc *ServiceClient) ActivateUrl(ctx *gin.Context) {
	routes.ActivateUrl(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserUrls(ctx *gin.Context) {
	routes.GetUserUrls(ctx, svc.Client)
}
