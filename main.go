package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/minhtri67/simplebank/api"
	"github.com/minhtri67/simplebank/db/sqlc"
)

const (
	dbDriver        = "postgres"
	dbSource        = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverIPAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverIPAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)

	}
}
