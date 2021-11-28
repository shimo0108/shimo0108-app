package models

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Calendar struct {
	Id         int       `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	Visibility bool      `json:"visibility" form:"visibility"`
	Color      string    `json:"color" form:"color"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
}

func CreateCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		cmd := `insert into calendars (name, visibility, color, created_at) values ($1, $2, $3, $4)`

		_, err = Db.Exec(cmd,
			c.FormValue("name"),
			true,
			c.FormValue("color"),
			time.Now())
		if err != nil {
			log.Fatalln(err)
		}
		ca := &Calendar{
			Name:       c.FormValue("name"),
			Visibility: true,
			Color:      c.FormValue("color"),
		}

		return c.JSON(http.StatusOK, ca)
	}
}

func GetCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {
		calendar := Calendar{}

		cmd := `select id, name, visibility, COALESCE(color,''), created_at
	          from calendars where id = $1`
		err = Db.QueryRow(cmd, c.Param("id")).Scan(
			&calendar.Id,
			&calendar.Name,
			&calendar.Visibility,
			&calendar.Color,
			&calendar.CreatedAt,
		)
		return c.JSON(http.StatusOK, &Calendar{Id: calendar.Id, Name: calendar.Name, Visibility: calendar.Visibility, Color: calendar.Color, CreatedAt: calendar.CreatedAt})
	}
}

func GetCalendars() echo.HandlerFunc {
	return func(c echo.Context) error {
		calendar := Calendar{}
		calendars := []*Calendar{}

		rows, err := Db.Query("select id, name, visibility, COALESCE(color,''), created_at from calendars")
		if err != nil {
			return errors.Wrapf(err, "cannot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(
				&calendar.Id,
				&calendar.Name,
				&calendar.Visibility,
				&calendar.Color,
				&calendar.CreatedAt); err != nil {
				return errors.Wrapf(err, "cannot connect SQL")
			}
			calendars = append(calendars, &Calendar{Id: calendar.Id, Name: calendar.Name, Visibility: calendar.Visibility, Color: calendar.Color, CreatedAt: calendar.CreatedAt})
		}

		return c.JSON(http.StatusOK, calendars)
	}
}

func UpdateCalendar() echo.HandlerFunc {
	return func(c echo.Context) error {

		cmd := `update calendars
						set name = $1
						,visibility = $2
						,color = $3
						where id = $4`

		_, err = Db.Exec(cmd,
			c.FormValue("name"),
			c.FormValue("visibility"),
			c.FormValue("color"),
			c.Param("id"))
		if err != nil {
			log.Fatalln(err)
		}
		e := &Calendar{
			Id:         stringToInt(c.Param("id")),
			Name:       c.FormValue("name"),
			Visibility: stringToBool(c.FormValue("visibility")),
			Color:      c.FormValue("color"),
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
