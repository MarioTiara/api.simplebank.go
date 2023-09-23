package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/marioTiara/api.simplebank.go/api"
	db "github.com/marioTiara/api.simplebank.go/db/sqlc"
	"github.com/marioTiara/api.simplebank.go/utils"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
