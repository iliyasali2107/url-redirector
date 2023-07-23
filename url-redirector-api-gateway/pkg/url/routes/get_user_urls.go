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

func GetUserURLs(ctx *gin.Context, client pb.URLServiceClient) {
	var reqBody GetUserUrlsRequestBody
	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect user id")
		return
	}

	req := &pb.GetUserURLsRequest{
		UserId: reqBody.UserId,
	}

	res, err := client.GetUserURLs(ctx, req)
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
