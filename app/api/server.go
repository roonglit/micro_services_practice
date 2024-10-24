package api

import db "github.com/JackleStyle0/micro_services_practice/db/sqlc"

type Server struct {
	store *db.Queries
}

func NewServer(queries *db.Queries) (*Server, error) {
	server := &Server{
		store: queries,
	}
	return server, nil
}
