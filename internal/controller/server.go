package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"math/rand"
	"time"
	"unicode/utf8"
	"url-services/internal/repo"
)

type Server struct {
	r repo.Storage
}

func NewServer(rp repo.Storage) *Server {
	return &Server{r: rp}
}

func (s *Server) Save(ctx context.Context, req *SaveRequest) (*SaveResponse, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	const keyLength = 10

	if req.GetOriginalUrl() == "" {
		return nil, fmt.Errorf("request not valid")
	}

	res, err := s.r.Check(ctx, req.GetOriginalUrl())
	if !errors.Is(err, pgx.ErrNoRows) && err != nil {
		log.Printf("Error Check: %s on %s", err, req.GetOriginalUrl())
		return nil, fmt.Errorf("error Check on %s", req.GetOriginalUrl())
	}
	if res != "" {
		return &SaveResponse{ShortUrl: res}, nil
	}

	rand.NewSource(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}

	err = s.r.Save(ctx, req.GetOriginalUrl(), string(shortKey))
	if err != nil {
		log.Printf("Error Save %s on %s", err, req.GetOriginalUrl())
		return nil, fmt.Errorf("error Save on %s", req.GetOriginalUrl())
	}

	return &SaveResponse{ShortUrl: string(shortKey)}, nil
}

func (s *Server) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	if utf8.RuneCountInString(req.GetShortUrl()) != 10 {
		return nil, fmt.Errorf("short url not valid")
	}
	res, err := s.r.Get(ctx, req.GetShortUrl())
	if err != nil {
		log.Printf("Error Get %s on %s", err, req.GetShortUrl())
		return nil, fmt.Errorf("error Get on %s", req.GetShortUrl())
	}

	return &GetResponse{OriginalUrl: res}, nil
}

func (s *Server) mustEmbedUnimplementedUrlServer() {
}
