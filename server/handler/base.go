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
	HOST     = "terraform-20211202131712468600000007.chsdgfmwm1wv.ap-northeast-1.rds.amazonaws.com"
	DATABASE = "shimo_app_db"
	USER     = "shimo0108"
	PASSWORD = "password"
	PORT     = "5432"
)

func init() {
	var connectionString string = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)
	Db, err = sql.Open("postgres", connectionString)
	fmt.Println(Db)

	if err != nil {
		log.Fatal(err)
	}
	_, err = Db.Exec("drop table calendars")
	if err != nil {
		fmt.Println(err)
	}

	_, err = Db.Exec("drop table events")
	if err != nil {
		fmt.Println(err)
	}

	calendars_cmd := `create table calendars (
		id          serial primary key,
		name        varchar(100) not null,
		visibility  boolean default true,
		color       varchar(16),
		created_at  timestamp not null default current_timestamp)`
	fmt.Println(calendars_cmd)

	_, err = Db.Exec(calendars_cmd)
	if err != nil {
		fmt.Println(err)
	}

	events_cmd := `create table events (
		id          serial primary key,
		name        varchar(100) not null,
		started_at  timestamp not null default current_timestamp,
		ended_at    timestamp not null default current_timestamp,
		calendar_id int,
		timed       boolean default true,
		description varchar(255),
		color       varchar(16),
		created_at  timestamp not null default current_timestamp,
		foreign key(calendar_id) references calendars(id))`
	fmt.Println(events_cmd)

	_, err = Db.Exec(events_cmd)
	if err != nil {
		fmt.Println(err)
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
