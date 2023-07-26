package routes

import (
	"net/http"

	"url-redirector-api-gateway/pkg/url/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SetActiveRequestBody struct {
	Id int64 `json:"id"`
}

func ActivateUrl(ctx *gin.Context, client pb.UrlServiceClient) {
	var reqBody SetActiveRequestBody
	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "id is incorrect form"+err.Error())
		return
	}
	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	req := &pb.ActivateUrlRequest{
		UrlId:  reqBody.Id,
		UserId: userId.(int64),
	}

	_, err = client.ActivateUrl(ctx, req)
	if err != nil {
		st, _ := status.FromError(err)
		switch st.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
		case codes.PermissionDenied:
			ctx.JSON(http.StatusForbidden, gin.H{"error": st.Message()})
		case codes.AlreadyExists:
			ctx.JSON(http.StatusForbidden, gin.H{"error": st.Message()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": st.Message()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "url is activated"})
}
