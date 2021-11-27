package models

import (
	"testing"
	"time"

	"github.com/labstack/echo"
)

type CalenderModesStub struct{}

func (c *CalenderModesStub) GetCalendars() ([]Calendar, error) {
	calendar := []Calendar{}
	t, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")

	c1 := Calendar{Name: "イベント", Visibility: true}
	c1.Id = 1
	c1.CreatedAt = t
	calendar = append(calendar, c1)

	c2 := Calendar{Name: "予定", Visibility: true}
	c2.Id = 2
	c2.CreatedAt = t
	calendar = append(calendar, c2)

	return calendar, nil
}

func TestCreateCalendar(t *testing.T) {
	// Setup
	e := echo.New()
	initRouting(e)
	cstub := &CalenderStub{}
	h := NewCalendarHandler(cstub)
	e.GET("/calendars", h.GetCalendars)
}
