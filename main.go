package main

import (
	"database/sql"
	"log"

	"github.com/LiShangAn/bank/api"
	db "github.com/LiShangAn/bank/db/sqlc"
	"github.com/LiShangAn/bank/util"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBdriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connecto to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
