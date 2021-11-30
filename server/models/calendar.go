package models

import (
	"database/sql"
	"log"
	"time"
)

type Calendar struct {
	Id         int       `json:"id" form:"id"`
	Name       string    `json:"name" form:"name"`
	Visibility bool      `json:"visibility" form:"visibility"`
	Color      string    `json:"color" form:"color"`
	CreatedAt  time.Time `json:"created_at" form:"created_at"`
}

func GetCalendar(id int, db *sql.DB) (calendar Calendar, err error) {
	calendar = Calendar{}
	cmd := `select id, name, visibility, COALESCE(color,''), created_at
	          from calendars where id = $1`
	err = db.QueryRow(cmd, id).Scan(
		&calendar.Id,
		&calendar.Name,
		&calendar.Visibility,
		&calendar.Color,
		&calendar.CreatedAt,
	)
	return calendar, err
}

func GetCalendars(db *sql.DB) (calendars []*Calendar, err error) {
	calendar := Calendar{}
	calendars = []*Calendar{}

	rows, err := db.Query("select id, name, visibility, COALESCE(color,''), created_at from calendars")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&calendar.Id,
			&calendar.Name,
			&calendar.Visibility,
			&calendar.Color,
			&calendar.CreatedAt); err != nil {
			panic("You can't open DB (dbGetAll())")
		}
		calendars = append(calendars, &Calendar{Id: calendar.Id, Name: calendar.Name, Visibility: calendar.Visibility, Color: calendar.Color, CreatedAt: calendar.CreatedAt})

	}
	return calendars, err

}

func (c *Calendar) CreateCalendar(db *sql.DB) (err error) {
	cmd := `insert into calendars (name, visibility, color, created_at) values ($1, $2, $3, $4)`

	_, err = db.Exec(cmd,
		c.Name,
		true,
		c.Color,
		time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (c *Calendar) UpdateCalendar(db *sql.DB) (err error) {
	cmd := `update calendars
						set name = $1
						,visibility = $2
						,color = $3
						where id = $4`

	_, err = db.Exec(cmd,
		c.Name,
		c.Visibility,
		c.Color,
		c.Id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func (c *Calendar) DeleteCalendar(db *sql.DB) (err error) {
	cmd_1 := `delete from events where calendar_id = $1`
	_, err = db.Exec(cmd_1, c.Id)
	if err != nil {
		log.Fatalln(err)
	}

	cmd_2 := `delete from calendars where id = $1`
	_, err = db.Exec(cmd_2, c.Id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
