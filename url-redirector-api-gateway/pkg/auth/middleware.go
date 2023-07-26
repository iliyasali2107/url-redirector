package auth

import (
	"context"
	"net/http"
	"strings"

	"url-redirector-api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
		return
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
		return
	}
	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "you are not authorized"})
		return
	}

	ctx.Set("userId", res.UserID)
	ctx.Next()
}
