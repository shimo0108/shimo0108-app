package handler

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	str_connect := ("host=postgres user=shimo0108 dbname=shimo_app_db password=password sslmode=disable")
	Db, err = sql.Open("postgres", str_connect)
	fmt.Println(Db)

	if err != nil {
		log.Fatal(err)
	}
	return
}

var layout = "2006-01-02 15:04:05"

func StringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func stringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func stringToBool(str string) bool {
	b, _ := strconv.ParseBool(str)
	return b
}
