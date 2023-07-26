package routes

import (
	"net/http"

	"url-redirector-api-gateway/pkg/url/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserUrlsRequestBody struct {
	UserId int64 `json:"user_id"`
}

func GetUserUrls(ctx *gin.Context, client pb.UrlServiceClient) {
	var reqBody GetUserUrlsRequestBody
	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect user id"})
		return
	}

	req := &pb.GetUserUrlsRequest{
		UserId: reqBody.UserId,
	}

	res, err := client.GetUserUrls(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			ctx.JSON(http.StatusNotFound, "you have no urls")
			return
		}

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
