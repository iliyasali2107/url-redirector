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
	pb.UnimplementedURLServiceServer
}

type Service interface {
	pb.URLServiceServer
	AddURL(context.Context, *pb.AddURLRequest) (*pb.AddURLResponse, error)
	GetURL(context.Context, *pb.GetURLRequest) (*pb.GetURLResponse, error)
	SetActiveURL(context.Context, *pb.SetActiveUrlRequest) (*pb.SetActiveUrlResponse, error)
	GetUserURLs(context.Context, *pb.GetUserURLsRequest) (*pb.GetUserURLsResponse, error)
}

func NewService(s db.Storage) Service {
	return &service{
		S: s,
	}
}

func (srv *service) AddURL(ctx context.Context, req *pb.AddURLRequest) (*pb.AddURLResponse, error) {
	url := models.URL{
		UserID: req.UserId,
		URL:    req.Url,
	}
	url, err := srv.S.InsertURL(url)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "something unexpected happened")
	}

	pbUrl := fromUrlToProto(url)

	return &pb.AddURLResponse{
		Status: http.StatusOK,
		Url:    pbUrl,
	}, nil
}

func (srv *service) GetURL(ctx context.Context, req *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	url, err := srv.S.GetActiveURL(req.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "there is no requested url")
		}

		return nil, status.Errorf(codes.Internal, "something unexpected happened")
	}

	return &pb.GetURLResponse{
		Status: http.StatusOK,
		Url:    fromUrlToProto(url),
	}, nil
}

func (srv *service) SetActiveURL(ctx context.Context, req *pb.SetActiveUrlRequest) (*pb.SetActiveUrlResponse, error) {
	activeURL, err := srv.S.GetActiveURL(req.UrlId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, err
	}

	if activeURL.UserID != req.UserId {
		return nil, status.Errorf(codes.PermissionDenied, "u are not allowed to activate this url")
	}

	if activeURL.ID == req.UrlId {
		return nil, status.Errorf(codes.AlreadyExists, "this url already activated")
	}

	_, err = srv.S.SetActive(req.UrlId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set active url: %v", err)
	}
	_, err = srv.S.SetNotActive(activeURL.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set unactive url: %v", err)
	}

	return &pb.SetActiveUrlResponse{
		Status: http.StatusOK,
	}, nil
}

func (srv *service) GetUserURLs(ctx context.Context, req *pb.GetUserURLsRequest) (*pb.GetUserURLsResponse, error) {
	urls, err := srv.S.GetUserURLs(req.UserId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.Internal, "failed to get user urls")
	}

	return &pb.GetUserURLsResponse{
		Status: http.StatusOK,
		Url:    fromUrlsToProtos(urls),
	}, nil
}

func fromUrlsToProtos(urls []models.URL) []*pb.URL {
	protoUrls := []*pb.URL{}
	for _, url := range urls {
		protoUrl := &pb.URL{
			Id:     url.ID,
			UserId: url.UserID,
			Url:    url.URL,
		}

		protoUrls = append(protoUrls, protoUrl)
	}
	return protoUrls
}

func fromUrlToProto(url models.URL) *pb.URL {
	return &pb.URL{
		Id:     url.ID,
		UserId: url.UserID,
		Url:    url.URL,
	}
}
