package main

import (
	"context"
	"log"

	"github.com/JackleStyle0/micro_services_practice/config"
	db "github.com/JackleStyle0/micro_services_practice/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource = "postgres://root:root@localhost:5432/auth_development?sslmode=disable"
)

var queries *db.Queries

func main() {

	conPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database")
	}

	queries = db.New(conPool)
	r := config.SetupRoutes()
	r.Run("localhost:3000")
}
