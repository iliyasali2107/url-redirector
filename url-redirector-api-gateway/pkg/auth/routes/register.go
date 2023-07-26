package routes

import (
	"context"
	"net/http"

	"url-redirector-api-gateway/pkg/auth/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,alpha,min=3"`
	Surname  string `json:"surname" binding:"required,alpha,min=3"`
	Password string `json:"password" binding:"required,alphanum,min=8"`
}

// @Summary Register new user
// @Tags auth
// @Description Register
// @Accept  json
// @Produce  json
// @Param input body RegisterRequestBody true "credentials"
// @Success 201 {object} pb.RegisterResponse
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /auth/register [post]

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	var req RegisterRequestBody

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid credentials"})
		return
	}
	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    req.Email,
		Name:     req.Name,
		Surname:  req.Surname,
		Password: req.Password,
	})
	if err != nil {
		status, _ := status.FromError(err)
		if status.Code() == codes.AlreadyExists {
			ctx.JSON(http.StatusConflict, "user already exists")
			return
		}

		ctx.JSON(http.StatusInternalServerError, "something unexpected happened"+err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
