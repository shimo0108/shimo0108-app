package handler

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

func init() {
	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")

		//もし err がnilではないなら、"読み込み出来ませんでした"が出力されます。
		if err != nil {
			fmt.Printf("読み込み出来ませんでした: %v", err)
		}
	}

	DBMS := "postgres"
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	var connection string = "host=" + DB_HOST + "\n" + "user=" + DB_USER + "\n" + "password=" + DB_PASS + "\n" + "dbname=" + DB_NAME + "\n" + "sslmode=disable"

	fmt.Println(connection)
	Db, err = sql.Open(DBMS, connection)

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
