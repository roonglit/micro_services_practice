package main

import (
	"context"
	"log"

	"github.com/JackleStyle0/micro_services_practice/app/api"
	db "github.com/JackleStyle0/micro_services_practice/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

var queries *db.Queries

const (
	dbSource = "postgres://root:root@localhost:5432/auth_development?sslmode=disable"
)

func main() {

	conPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	queries := db.New(conPool)
	server, _ := api.NewServer(queries)
	r := server.SetupRoutes()
	r.Run("localhost:3000")
}
