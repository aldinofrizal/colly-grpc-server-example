package services

import (
	"context"

	pb "github.com/aldinofrizal/soompi-colly-scrap/news"
)

type Server struct {
	pb.UnimplementedNewsServiceServer
}

func (s *Server) GetNews(ctx context.Context, in *pb.EmptyParams) (*pb.ItemResponse, error) {
	items, err := GetNews()
	if err != nil {
		return nil, err
	}
	itemResponse := pb.ItemResponse{
		Items: items,
	}

	return &itemResponse, nil
}
