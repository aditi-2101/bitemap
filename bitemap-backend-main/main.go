package main

import (
	"context"
	"log"

	"bitemap/api"
	db "bitemap/db/sqlc"
	"bitemap/util"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuation: ", err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
