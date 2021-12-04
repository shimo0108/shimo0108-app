package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	USER := os.Getenv("RDS_USER")
	PASS := os.Getenv("RDS_PASSWORD")
	PROTOCOL := "tcp(" + os.Getenv("RDS_HOST") + ":5432)"
	DBNAME := os.Getenv("RDS_DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "sslmode=disable"

	Db, err = sql.Open("postgres", CONNECT)
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
