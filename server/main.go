package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shimo0108/shimo0108-app/server/handler"
)

func main() {
	router := NewRouter()

	router.Start(":9999")
}

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	initRouting(e)

	return e
}

func initRouting(e *echo.Echo) {
	e.GET("/hello", helloHandler)
	e.GET("/api/v1/events", handler.GetEventsHandler())
	e.POST("/api/v1/events", handler.CreateEventHandler())
	e.PUT("/api/v1/events/:id", handler.UpdateEventHandler())
	e.DELETE("/api/v1/events/:id", handler.DeleteEventHandler())

	e.GET("/api/v1/calendars/:id", handler.GetCalendarHandler())
	e.GET("/api/v1/calendars", handler.GetCalendarsHandler())
	e.POST("/api/v1/calendars", handler.CreateCalendarHandler())
	e.PUT("/api/v1/calendars/:id", handler.UpdateCalendarHandler())
	e.DELETE("/api/v1/calendars/:id", handler.DeleteCalendarHandler())
}

func helloHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo World!!")
}
