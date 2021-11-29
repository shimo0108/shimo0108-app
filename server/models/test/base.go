package test

import (
	"database/sql/driver"
	"strconv"
	"time"
)

type AnyTime struct{}

var layout = "2006-01-02 15:04:05"

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

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
