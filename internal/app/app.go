package app

import (
	"context"
	"log"
	"url-services/internal/repo"
	in_memory "url-services/pkg/in-memory"
	"url-services/pkg/postgres"
)

func Run(param string) {
	log.Println("Url-services start work...")
	ctx := context.Background()

	switch param {
	case "memory":
		db := in_memory.NewMemory()
		rp := repo.NewMemoRepo(db)
	case "pst":
		db := postgres.NewPostgres(ctx)
		rp := repo.NewDbRepo(db)
	default:
		log.Fatal("Error, params db not correct")
	}
}
