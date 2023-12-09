package controller

import (
	"context"
	"testing"
	"url-services/internal/repo"
	"url-services/pkg/postgres"
)

func Test_server_Save(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "url-services")
	t.Setenv("pst_port", "5430")
	ctx := context.Background()
	sr := NewServer(repo.NewDbRepo(postgres.NewPostgres(ctx)))

	tests := []struct {
		name string
		req  string
		resp string
	}{
		{
			name: "test1",
			req:  "https://www.google.com/",
			resp: "qwertyuiop",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := sr.Save(context.Background(), &SaveRequest{OriginalUrl: tt.req})
			if err != nil || resp.GetShortUrl() != tt.resp {
				t.Error("Save server error")
				return
			}
		})
	}
}

func Test_server_Get(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "url-services")
	t.Setenv("pst_port", "5430")
	ctx := context.Background()
	sr := NewServer(repo.NewDbRepo(postgres.NewPostgres(ctx)))

	tests := []struct {
		name string
		req  string
		resp string
	}{
		{
			name: "test1",
			req:  "qwertyuiop",
			resp: "https://www.google.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := sr.Get(context.Background(), &GetRequest{ShortUrl: tt.req})
			if err != nil || resp.GetOriginalUrl() != tt.resp {
				t.Error("Get server error")
				return
			}
		})
	}
}
