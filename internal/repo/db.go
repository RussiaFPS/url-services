package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Storage interface {
	Save(ctx context.Context, urlOriginal string, urlShort string) error
	Get(ctx context.Context, urlShort string) (string, error)
}

type DbRepo struct {
	db *pgxpool.Pool
}

type MemoRepo struct {
	db *sync.Map
}

func NewDbRepo(pg *pgxpool.Pool) Storage {
	return &DbRepo{db: pg}
}

func NewMemoRepo(memo *sync.Map) Storage {
	return &MemoRepo{db: memo}
}

func (d *DbRepo) Save(ctx context.Context, urlShort string, urlOriginal string) error {
	q := `insert into urls (url_shorts,url_original) VALUES ($1,$2)`

	if _, err := d.db.Exec(ctx, q, urlShort, urlOriginal); err != nil {
		return err
	}

	return nil
}

func (d *DbRepo) Get(ctx context.Context, urlShort string) (string, error) {
	var urlOriginal string
	q := `select url_original from urls where url_shorts=$1`

	if err := d.db.QueryRow(ctx, q, urlShort).Scan(&urlOriginal); err != nil {
		return "", err
	}

	return urlOriginal, nil
}

func (m *MemoRepo) Save(ctx context.Context, urlOriginal string, urlShort string) error {
	return nil
}
func (m *MemoRepo) Get(ctx context.Context, urlShort string) (string, error) {
	return "", nil
}
