package routes

import (
	"context"
	"net/http"
	"strconv"

	"url-redirector-api-gateway/pkg/url/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUrl(ctx *gin.Context, client pb.UrlServiceClient) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, "incorrect id")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "incorrect id")
		return
	}

	request := &pb.GetUrlRequest{Id: id}

	res, err := client.GetUrl(context.Background(), request)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": st.Message()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusSeeOther, res.Url.Url)
}
