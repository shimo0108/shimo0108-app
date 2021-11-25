package models

import (
	"log"
	"time"
)

type Calendar struct {
	Id         int
	Name       string
	Visibility bool
	CreatedAt  time.Time
}

func (c *Calendar) CreateCalendar() (err error) {
	cmd := `insert into calendars (name, visibility, created_at) values ($1, $2, $3)`

	_, err = Db.Exec(cmd,
		c.Name,
		c.Visibility,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetCalendar(id int) (calendar Calendar, err error) {
	calendar = Calendar{}
	cmd := `select id, name, visibility, created_at
	from calendars where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&calendar.Id,
		&calendar.Name,
		&calendar.Visibility,
		&calendar.CreatedAt,
	)
	return calendar, err
}

func (c *Calendar) UpdateCalendar() (err error) {
	cmd := `update calendars
	set name = $1
	,visibility = $2
  where id = $3`
	_, err = Db.Exec(cmd, c.Name, c.Visibility, c.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (c *Calendar) DeleteCalendar() (err error) {
	cmd := `delete from calendars where id = $1`
	_, err = Db.Exec(cmd, c.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
