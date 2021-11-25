package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	str_connect := ("host=postgres user=shimo0108 dbname=task_list_db password=password sslmode=disable")
	Db, err = sql.Open("postgres", str_connect)
	fmt.Println(Db)

	if err != nil {
		log.Fatal(err)
	}
	return
}
