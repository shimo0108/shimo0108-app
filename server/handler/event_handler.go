package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/shimo0108/shimo0108-app/server/models"
)

func GetEventsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		Events, err := models.GetEvents(Db)
		if err != nil {
			log.Fatalln(err)
		}

		return c.JSON(http.StatusOK, Events)
	}
}

func CreateEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := &models.Event{}
		e.Name = c.FormValue("name")
		e.StartTime = StringToTime(c.FormValue("start"))
		e.EndTime = StringToTime(c.FormValue("end"))
		e.CalendarId = stringToInt(c.FormValue("calendar_id"))
		e.Timed = stringToBool(c.FormValue("timed"))
		e.Description = c.FormValue("description")
		e.Color = c.FormValue("color")

		e.CreateEvent(Db)

		return c.JSON(http.StatusOK, e)
	}
}

func UpdateEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := &models.Event{}
		e.Id = stringToInt(c.Param("id"))
		e.Name = c.FormValue("name")
		e.StartTime = StringToTime(c.FormValue("start"))
		e.EndTime = StringToTime(c.FormValue("end"))
		e.CalendarId = stringToInt(c.FormValue("calendar_id"))
		e.Timed = stringToBool(c.FormValue("timed"))
		e.Description = c.FormValue("description")
		e.Color = c.FormValue("color")
		e.UpdateEvent(Db)

		return c.JSON(http.StatusOK, e)
	}
}

func DeleteEventHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		e := &models.Event{}
		e.Id = stringToInt(c.Param("id"))
		e.DeleteEvent(Db)

		return c.JSON(http.StatusOK, e)
	}
}
