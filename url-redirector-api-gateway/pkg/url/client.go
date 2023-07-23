package url

import (
	"fmt"
	"url-redirector-api-gateway/pkg/config"
	"url-redirector-api-gateway/pkg/url/pb"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.URLServiceClient
}

func InitServiceClient(c *config.Config) pb.URLServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.UrlSvcPort, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewURLServiceClient(cc)
}
