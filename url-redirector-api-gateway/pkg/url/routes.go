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

	r.GET("/urls/:id", svc.GetURL)
	r.Use(a.AuthRequired)
	r.GET("/urls", svc.GetUserURLs)
	r.POST("/urls", svc.AddURL)
	r.POST("/urls/activate", svc.SetActiveURL)

	return svc
}

func (svc *ServiceClient) GetURL(ctx *gin.Context) {
	routes.GetURL(ctx, svc.Client)
}

func (svc *ServiceClient) AddURL(ctx *gin.Context) {
	routes.AddURL(ctx, svc.Client)
}

func (svc *ServiceClient) SetActiveURL(ctx *gin.Context) {
	routes.SetActiveURL(ctx, svc.Client)
}

func (svc *ServiceClient) GetUserURLs(ctx *gin.Context) {
	routes.GetUserURLs(ctx, svc.Client)
}
