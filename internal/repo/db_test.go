package repo

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"testing"
	in_memory "url-services/pkg/in-memory"
	"url-services/pkg/postgres"
)

func Test_db_Save(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "url-services")
	t.Setenv("pst_port", "5430")
	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name        string
		urlOriginal string
		urlShort    string
	}{
		{
			name:        "test1",
			urlOriginal: "https://www.google.com/",
			urlShort:    "qwertyuiop",
		},
		{
			name:        "test2",
			urlOriginal: "https://ya.ru/",
			urlShort:    "zxcvbnmasd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			err := d.Save(context.Background(), tt.urlOriginal, tt.urlShort)
			if err != nil {
				t.Error("Save db error")
				return
			}
		})
	}
}

func Test_db_Get(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "url-services")
	t.Setenv("pst_port", "5430")
	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name     string
		urlShort string
		want     string
		err      error
	}{
		{
			name:     "test1",
			urlShort: "qwertyuiop",
			want:     "https://www.google.com/",
			err:      nil,
		},
		{
			name:     "test2",
			urlShort: "test",
			want:     "",
			err:      pgx.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			res, err := d.Get(context.Background(), tt.urlShort)
			if res != tt.want || !errors.Is(err, tt.err) {
				t.Error("Get db error")
				return
			}
		})
	}
}

func Test_db_Check(t *testing.T) {
	t.Setenv("pst_host", "localhost")
	t.Setenv("pst_user", "postgres")
	t.Setenv("pst_password", "qwer1234")
	t.Setenv("pst_dbname", "url-services")
	t.Setenv("pst_port", "5430")
	pg := postgres.NewPostgres(context.Background())

	tests := []struct {
		name        string
		urlOriginal string
		want        string
		err         error
	}{
		{
			name:        "test1",
			urlOriginal: "https://www.google.com/",
			want:        "qwertyuiop",
			err:         nil,
		},
		{
			name:        "test2",
			urlOriginal: "ya",
			want:        "",
			err:         pgx.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DbRepo{
				db: pg,
			}

			res, err := d.Check(context.Background(), tt.urlOriginal)
			if res != tt.want || !errors.Is(err, tt.err) {
				t.Error("Check db error")
				return
			}
		})
	}
}

func Test_memo_Save(t *testing.T) {
	memo := in_memory.NewMemory()

	tests := []struct {
		name        string
		urlOriginal string
		urlShort    string
	}{
		{
			name:        "test1",
			urlOriginal: "https://www.google.com/",
			urlShort:    "qwertyuiop",
		},
		{
			name:        "test2",
			urlOriginal: "https://ya.ru/",
			urlShort:    "zxcvbnmasd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MemoRepo{
				db: memo,
			}

			err := d.Save(context.Background(), tt.urlOriginal, tt.urlShort)
			if err != nil {
				t.Error("Save memo error")
				return
			}
		})
	}
}

func Test_memo_Get(t *testing.T) {
	memo := in_memory.NewMemory()
	memo.Store("qwertyuiop", "https://www.google.com/")
	memo.Store("zxcvbnmasd", "https://ya.ru/")

	tests := []struct {
		name     string
		urlShort string
		want     string
		err      error
	}{
		{
			name:     "test1",
			urlShort: "qwertyuiop",
			want:     "https://www.google.com/",
			err:      nil,
		},
		{
			name:     "test2",
			urlShort: "test",
			want:     "",
			err:      pgx.ErrNoRows,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MemoRepo{
				db: memo,
			}

			res, err := d.Get(context.Background(), tt.urlShort)
			if res != tt.want || !errors.Is(err, tt.err) {
				t.Error("Get memo error")
				return
			}
		})
	}
}

func Test_memo_Check(t *testing.T) {
	memo := in_memory.NewMemory()
	memo.Store("qwertyuiop", "https://www.google.com/")
	memo.Store("zxcvbnmasd", "https://ya.ru/")

	tests := []struct {
		name        string
		urlOriginal string
		want        string
	}{
		{
			name:        "test1",
			urlOriginal: "https://www.google.com/",
			want:        "qwertyuiop",
		},
		{
			name:        "test2",
			urlOriginal: "ya",
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &MemoRepo{
				db: memo,
			}

			res, err := d.Check(context.Background(), tt.urlOriginal)
			if res != tt.want || err != nil {
				t.Error("Check memo error")
				return
			}
		})
	}
}
