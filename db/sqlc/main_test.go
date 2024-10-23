package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	// dbDriver = "postgresql"
	dbSource = "postgres://root:root@localhost:5432/auth_development?sslmode=disable"
)

func TestMain(m *testing.M) {

	conPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("can't connect db", err)
	}
	testQueries = New(conPool)

	os.Exit(m.Run())

}
