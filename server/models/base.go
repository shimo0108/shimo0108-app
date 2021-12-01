package models

import (
	"strconv"
	"time"
)

var layout = "2006-01-02 15:04:05"

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

func stringToTime(str string) time.Time {
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
