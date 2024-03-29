package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/minhtri67/simplebank/api"
	"github.com/minhtri67/simplebank/db/sqlc"
	"github.com/minhtri67/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlc.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	err = server.Start(config.ServerIPAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
