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

const (
	HOST     = "terraform-20211202131712468600000007.chsdgfmwm1wv.ap-northeast-1.rds.amazonaws.com:5432"
	DATABASE = "shimo_app_db"
	USER     = "shimo0108"
	PASSWORD = "password"
)

type Sale struct {
	Loginid  string
	Password string
}

func init() {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	Db, err = sql.Open("postgres", connectionString)
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
