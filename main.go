package main

import (
	"database/sql"
	"log"

	"github.com/umeh-promise/simple_bank/api"
	db "github.com/umeh-promise/simple_bank/db/sqlc"
	"github.com/umeh-promise/simple_bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load configurations", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to the database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
