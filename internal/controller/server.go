package controller

import (
	"context"
	"url-services/internal/repo"
)

type Server struct {
	r repo.Storage
}

func NewServer(rp repo.Storage) *Server {
	return &Server{r: rp}
}

func (s *Server) Save(context.Context, *SaveRequest) (*SaveResponse, error) {
	return &SaveResponse{ShortUrl: "123"}, nil
}

func (s *Server) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return &GetResponse{OriginalUrl: "123"}, nil
}

func (s *Server) mustEmbedUnimplementedUrlServer() {
}
