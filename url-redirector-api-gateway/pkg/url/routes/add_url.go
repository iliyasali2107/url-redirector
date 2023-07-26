package routes

import (
	"fmt"
	"net/http"
	"strings"

	"url-redirector-api-gateway/pkg/url/pb"

	"github.com/gin-gonic/gin"
)

type AddUrlRequestBody struct {
	Url string `json:"url"`
}

func AddUrl(ctx *gin.Context, client pb.UrlServiceClient) {
	var reqBody AddUrlRequestBody

	err := ctx.BindJSON(&reqBody)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "url is not correct")
		return
	}

	if !strings.HasPrefix(reqBody.Url, "https://") {
		reqBody.Url = "https://" + reqBody.Url
	}

	userId, ok := ctx.Get("userId")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, nil)
		return
	}

	req := &pb.AddUrlRequest{
		UserId: userId.(int64),
		Url:    reqBody.Url,
	}

	fmt.Println(req)

	res, err := client.AddUrl(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("%s was added to url list", res.Url.Url))
}
