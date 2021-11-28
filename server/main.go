package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shimo0108/shimo0108-app/server/models"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	initRouting(e)

	e.Start(":9999")
}

func initRouting(e *echo.Echo) {
	e.GET("/api/v1/events", models.GetEvents())
	e.POST("/api/v1/events", models.CreateEvent())
	e.PUT("/api/v1/events/:id", models.UpdateEvent())
	e.DELETE("/api/v1/events/:id", models.DeleteEvent())

	e.GET("/api/v1/calendars/:id", models.GetCalendar())
	e.GET("/api/v1/calendars", models.GetCalendars())
	e.POST("/api/v1/calendars", models.CreateCalendar())
	e.PUT("/api/v1/calendars/:id", models.UpdateCalendar())
	e.DELETE("/api/v1/calendars/:id", models.DeleteCalendar())
}
