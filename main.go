package main

import (
	"log"

	"github.com/gopheramit/simple_bank/api"
	db "github.com/gopheramit/simple_bank/db"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := db.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
