package models

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Event struct {
	Id          int       `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	StartTime   time.Time `json:"start_time" form:"start"`
	EndTime     time.Time `json:"end_time" form:"end"`
	CalendarId  int       `json:"calendar_id" form:"calendar_id"`
	Timed       bool      `json:"timed" form:"timed"`
	Description string    `json:"description" form:"description"`
	Color       string    `json:"color" form:"color"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}

func CreateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := `insert into events (name, start_time, end_time, calendar_id, timed, description, color, created_at) values ($1, $2, $3, $4, $5, $6 ,$7, $8)`

		_, err = Db.Exec(cmd,
			c.FormValue("name"),
			c.FormValue("start"),
			c.FormValue("end"),
			c.FormValue("calendar_id"),
			c.FormValue("timed"),
			c.FormValue("description"),
			c.FormValue("color"),
			time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		e := &Event{
			Name:        c.FormValue("name"),
			StartTime:   stringToTime(c.FormValue("start")),
			EndTime:     stringToTime(c.FormValue("end")),
			CalendarId:  stringToInt(c.FormValue("calendar_id")),
			Timed:       stringToBool(c.FormValue("timed")),
			Description: c.FormValue("description"),
			Color:       c.FormValue("color"),
		}

		return c.JSON(http.StatusOK, e)
	}
}

func GetEvents() echo.HandlerFunc {
	return func(c echo.Context) error {
		event := Event{}
		events := []*Event{}

		rows, err := Db.Query("select id, name, start_time, end_time, calendar_id, timed, COALESCE(description,''), COALESCE(color,''), created_at from events")
		if err != nil {
			return errors.Wrapf(err, "cannot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(
				&event.Id,
				&event.Name,
				&event.StartTime,
				&event.EndTime,
				&event.CalendarId,
				&event.Timed,
				&event.Description,
				&event.Color,
				&event.CreatedAt); err != nil {
				return errors.Wrapf(err, "cannot connect SQL")
			}
			events = append(events, &Event{Id: event.Id, Name: event.Name, StartTime: event.StartTime, EndTime: event.EndTime, CalendarId: event.CalendarId, Timed: event.Timed, Description: event.Description, Color: event.Color, CreatedAt: event.CreatedAt})
		}

		return c.JSON(http.StatusOK, events)
	}
}

func UpdateEvent() echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := `update events set name = $1 ,start_time = $2 ,end_time = $3 ,calendar_id = $4 ,timed = $5 ,description = $6 ,color = $7 where id = $8`

		_, err = Db.Exec(cmd,
			c.FormValue("name"),
			c.FormValue("start_time"),
			c.FormValue("end_time"),
			c.FormValue("calendar_id"),
			c.FormValue("timed"),
			c.FormValue("description"),
			c.FormValue("color"),
			c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}
		e := &Event{
			Id:          stringToInt(c.Param("id")),
			Name:        c.FormValue("name"),
			StartTime:   stringToTime(c.FormValue("start_time")),
			EndTime:     stringToTime(c.FormValue("end_time")),
			CalendarId:  stringToInt(c.FormValue("calendar_id")),
			Timed:       stringToBool(c.FormValue("timed")),
			Description: c.FormValue("description"),
			Color:       c.FormValue("color"),
		}

		return c.JSON(http.StatusOK, e)
	}
}

func DeleteEvent() echo.HandlerFunc {
	return func(c echo.Context) error {
		cmd := `delete from events where id = $1`
		_, err = Db.Exec(cmd, c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}

		return c.JSON(http.StatusOK, "success")
	}
}
