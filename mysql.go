package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3307)/golang_api_learn")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
