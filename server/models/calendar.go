package models

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Calendar struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Visibility bool      `json:"visibility,string"`
	CreatedAt  time.Time `json:"created_at"`
}

func CreateCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := `insert into calendars (name, visibility, created_at) values ($1, $2, $3)`

		_, err = Db.Exec(cmd,
			c.QueryParam("name"),
			c.QueryParam("visibility"),
			time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		ca := &Calendar{
			Name:       c.QueryParam("name"),
			Visibility: stringToBool(c.QueryParam("visibility")),
		}

		return c.JSON(http.StatusOK, ca)
	}
}

func GetCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		calendar := Calendar{}

		cmd := `select id, name, visibility ,created_at
	          from calendars where id = $1`
		err = Db.QueryRow(cmd, c.Param("id")).Scan(
			&calendar.Id,
			&calendar.Name,
			&calendar.Visibility,
			&calendar.CreatedAt,
		)
		return c.JSON(http.StatusOK, &Calendar{Id: calendar.Id, Name: calendar.Name, Visibility: calendar.Visibility, CreatedAt: calendar.CreatedAt})
	}
}

func GetCalendars() echo.HandlerFunc {
	return func(c echo.Context) error {
		calendar := Calendar{}
		calendars := []*Calendar{}

		rows, err := Db.Query("select id, name, visibility, created_at from calendars")
		if err != nil {
			return errors.Wrapf(err, "cannot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(
				&calendar.Id,
				&calendar.Name,
				&calendar.Visibility,
				&calendar.CreatedAt); err != nil {
				return errors.Wrapf(err, "cannot connect SQL")
			}
			calendars = append(calendars, &Calendar{Id: calendar.Id, Name: calendar.Name, Visibility: calendar.Visibility, CreatedAt: calendar.CreatedAt})
		}

		return c.JSON(http.StatusOK, calendars)
	}
}

func UpdateCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := `update calendars
						set name = $1
						,visibility = $2
						where id = $3`

		_, err = Db.Exec(cmd,
			c.QueryParam("name"),
			c.QueryParam("visibility"),
			c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}
		e := &Calendar{
			Id:         stringToInt(c.Param("id")),
			Name:       c.QueryParam("name"),
			Visibility: stringToBool(c.QueryParam("visibility")),
		}

		return c.JSON(http.StatusOK, e)
	}
}

func DeleteCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		cmd_1 := `delete from events where calendar_id = $1`
		_, err = Db.Exec(cmd_1, c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}

		cmd_2 := `delete from calendars where id = $1`
		_, err = Db.Exec(cmd_2, c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}

		return c.JSON(http.StatusOK, "success")
	}
}
