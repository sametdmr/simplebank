package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sametdmr/simplebank/api"
	db "github.com/sametdmr/simplebank/db/sqlc"
)

const (
	dbDrive       = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}