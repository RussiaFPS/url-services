package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Storage interface {
	Save(ctx context.Context, urlOriginal string, urlShort string) error
	Get(ctx context.Context, urlShort string) (string, error)
	Check(ctx context.Context, urlOriginal string) (string, error)
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

func (d *DbRepo) Save(ctx context.Context, urlOriginal string, urlShort string) error {
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

func (d *DbRepo) Check(ctx context.Context, urlOriginal string) (string, error) {
	var urlShort string
	q := `select url_shorts from urls where url_original=$1`

	if err := d.db.QueryRow(ctx, q, urlOriginal).Scan(&urlShort); err != nil {
		return "", err
	}

	return urlShort, nil
}

func (m *MemoRepo) Save(ctx context.Context, urlOriginal string, urlShort string) error {
	m.db.Store(urlShort, urlOriginal)
	return nil
}
func (m *MemoRepo) Get(ctx context.Context, urlShort string) (string, error) {
	v, ok := m.db.Load(urlShort)
	if !ok {
		return "", fmt.Errorf("not found")
	}

	return v.(string), nil
}

func (m *MemoRepo) Check(ctx context.Context, urlOriginal string) (string, error) {
	var urlShort string

	m.db.Range(func(k, v interface{}) bool {
		if v == urlOriginal {
			urlShort = k.(string)
			return true
		}
		return false
	})

	return urlShort, nil
}
