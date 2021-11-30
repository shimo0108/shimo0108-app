package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/shimo0108/shimo0108-app/server/models"
)

func GetCalendarHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id int
		id, _ = strconv.Atoi(c.Param("id"))
		ca, _ := models.GetCalendar(id, Db)

		return c.JSON(http.StatusOK, &models.Calendar{Id: ca.Id, Name: ca.Name, Visibility: ca.Visibility, Color: ca.Color, CreatedAt: ca.CreatedAt})
	}
}

func GetCalendarsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		calendars, _ := models.GetCalendars(Db)

		return c.JSON(http.StatusOK, calendars)
	}
}

func CreateCalendarHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ca := &models.Calendar{}
		ca.Name = c.FormValue("name")
		ca.Color = c.FormValue("color")
		ca.CreateCalendar(Db)

		return c.JSON(http.StatusOK, ca)
	}
}

func UpdateCalendarHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ca := &models.Calendar{}
		ca.Id = stringToInt(c.Param("id"))
		ca.Name = c.FormValue("name")
		ca.Visibility = stringToBool(c.FormValue("visibility"))
		ca.Color = c.FormValue("color")
		ca.UpdateCalendar(Db)

		return c.JSON(http.StatusOK, ca)
	}
}

func DeleteCalendarHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		ca := &models.Calendar{}
		ca.Id = stringToInt(c.Param("id"))
		ca.DeleteCalendar(Db)

		return c.JSON(http.StatusOK, ca)
	}
}
