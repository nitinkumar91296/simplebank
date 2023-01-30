package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nitinkumar91296/simplebank/api"
	db "github.com/nitinkumar91296/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	var err error
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(serverAddress); err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
