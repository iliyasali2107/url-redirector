package routes

import (
	"context"
	"net/http"

	"url-redirector-api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,alphanum"`
}

// @Summary Login existing user
// @Tags auth
// @Description login
// @Accept  json
// @Produce  json
// @Param input body LoginRequestBody true "credentials"
// @Success 200 {object} pb.LoginResponse
// @Failure 404
// @Failure 400
// @Failure 500
// @Router /auth/login [post]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	var req LoginRequestBody

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		status, _ := status.FromError(err)
		switch status.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, status.Message())
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, status.Message())
		default:
			ctx.JSON(http.StatusInternalServerError, "something unexpected occured")
		}

		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
