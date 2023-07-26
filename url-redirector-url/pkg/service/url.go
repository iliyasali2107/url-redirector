package service

import (
	"context"
	"net/http"

	"url-redirecter-url/pkg/db"
	"url-redirecter-url/pkg/models"
	"url-redirecter-url/pkg/pb"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
	S db.Storage
	pb.UnimplementedUrlServiceServer
}

type Service interface {
	pb.UrlServiceServer
	AddUrl(context.Context, *pb.AddUrlRequest) (*pb.AddUrlResponse, error)
	GetUrl(context.Context, *pb.GetUrlRequest) (*pb.GetUrlResponse, error)
	ActivateUrl(context.Context, *pb.ActivateUrlRequest) (*pb.ActivateUrlResponse, error)
	GetUserUrls(context.Context, *pb.GetUserUrlsRequest) (*pb.GetUserUrlsResponse, error)
}

func NewService(s db.Storage) Service {
	return &service{
		S: s,
	}
}

func (srv *service) AddUrl(ctx context.Context, req *pb.AddUrlRequest) (*pb.AddUrlResponse, error) {
	url := models.Url{
		UserID: req.UserId,
		Url:    req.Url,
	}
	url, err := srv.S.InsertUrl(url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something unexpected happened")
	}

	pbUrl := fromUrlToProto(url)

	return &pb.AddUrlResponse{
		Status: http.StatusOK,
		Url:    pbUrl,
	}, nil
}

func (srv *service) GetUrl(ctx context.Context, req *pb.GetUrlRequest) (*pb.GetUrlResponse, error) {
	url, err := srv.S.GetActiveUrl(req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "there is no active url")
		}

		return nil, status.Errorf(codes.Internal, "something unexpected happened")
	}

	return &pb.GetUrlResponse{
		Status: http.StatusOK,
		Url:    fromUrlToProto(url),
	}, nil
}

func (srv *service) ActivateUrl(ctx context.Context, req *pb.ActivateUrlRequest) (*pb.ActivateUrlResponse, error) {
	dbUrl, err := srv.S.GetUrl(req.UrlId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "there is no requested url")
	}

	if dbUrl.UserID != req.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "u are not allowed to activate this url")
	}

	if dbUrl.Active {
		return nil, status.Errorf(codes.AlreadyExists, "this url is already activated")
	}

	activeUrl, err := srv.S.GetActiveUrl(req.UserId)
	if err == nil {
		_, err := srv.S.Deactivate(activeUrl.ID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to deactivet url: %v", err)
		}
	}

	_, err = srv.S.Activate(req.UrlId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to activate url: %v", err)
	}

	return &pb.ActivateUrlResponse{
		Status: http.StatusOK,
	}, nil
}

func (srv *service) GetUserUrls(ctx context.Context, req *pb.GetUserUrlsRequest) (*pb.GetUserUrlsResponse, error) {
	urls, err := srv.S.GetUserUrls(req.UserId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to get user urls")
	}

	return &pb.GetUserUrlsResponse{
		Status: http.StatusOK,
		Urls:   fromUrlsToProtos(urls),
	}, nil
}

func fromUrlsToProtos(urls []models.Url) []*pb.Url {
	protoUrls := []*pb.Url{}
	for _, url := range urls {
		protoUrl := &pb.Url{
			Id:     url.ID,
			UserId: url.UserID,
			Url:    url.Url,
			Active: url.Active,
		}

		protoUrls = append(protoUrls, protoUrl)
	}
	return protoUrls
}

func fromUrlToProto(url models.Url) *pb.Url {
	return &pb.Url{
		Id:     url.ID,
		UserId: url.UserID,
		Url:    url.Url,
		Active: url.Active,
	}
}
