package app

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"url-services/internal/controller"
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

		s := grpc.NewServer()
		srv := controller.NewServer(rp)
		controller.RegisterUrlServer(s, srv)

		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal("Error,listen 8080: ", err)
		}
		if err = s.Serve(l); err != nil {
			log.Fatal("Error,Serve: ", err)
		}
	case "pst":
		db := postgres.NewPostgres(ctx)
		rp := repo.NewDbRepo(db)

		s := grpc.NewServer()
		srv := controller.NewServer(rp)
		controller.RegisterUrlServer(s, srv)

		l, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal("Error,listen 8080: ", err)
		}
		if err = s.Serve(l); err != nil {
			log.Fatal("Error,Serve: ", err)
		}
	default:
		log.Fatal("Error, params db not correct")
	}
}
