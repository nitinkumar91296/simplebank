package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nitinkumar91296/simplebank/api"
	db "github.com/nitinkumar91296/simplebank/db/sqlc"
	"github.com/nitinkumar91296/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
