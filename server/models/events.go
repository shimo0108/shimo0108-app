package models

import (
	"database/sql"
	"log"
	"time"
)

type Event struct {
	Id          int       `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	StartedAt   time.Time `json:"started_at" form:"start"`
	EndedAt     time.Time `json:"ended_at" form:"end"`
	CalendarId  int       `json:"calendar_id" form:"calendar_id"`
	Timed       bool      `json:"timed" form:"timed"`
	Description string    `json:"description" form:"description"`
	Color       string    `json:"color" form:"color"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
}

func (e *Event) CreateEvent(db *sql.DB) (id int, err error) {
	cmd := `insert into events (name, started_at, ended_at, calendar_id, timed, description, color, created_at) values ($1, $2, $3, $4, $5, $6 ,$7, $8) RETURNING id`

	err = db.QueryRow(cmd,
		e.Name,
		e.StartedAt,
		e.EndedAt,
		e.CalendarId,
		e.Timed,
		e.Description,
		e.Color,
		time.Now()).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}

	return id, err
}

func (e *Event) UpdateEvent(db *sql.DB) (err error) {
	cmd := `update events set name = $1 ,started_at = $2 ,ended_at = $3 ,calendar_id = $4 ,timed = $5 ,description = $6 ,color = $7 where id = $8`

	_, err = db.Exec(cmd,
		e.Name,
		e.StartedAt,
		e.EndedAt,
		e.CalendarId,
		e.Timed,
		e.Description,
		e.Color,
		e.Id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func GetEvents(db *sql.DB) (events []*Event, err error) {
	event := Event{}
	events = []*Event{}

	rows, err := db.Query("select id, name, started_at, ended_at, calendar_id, timed, COALESCE(description,''), COALESCE(color,''), created_at from events")
	if err != nil {
		panic("You can't open DB (dbGetAll())")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&event.Id,
			&event.Name,
			&event.StartedAt,
			&event.EndedAt,
			&event.CalendarId,
			&event.Timed,
			&event.Description,
			&event.Color,
			&event.CreatedAt); err != nil {
			panic("You can't open DB (dbGetAll())")
		}
		events = append(events, &Event{Id: event.Id, Name: event.Name, StartedAt: event.StartedAt, EndedAt: event.EndedAt, CalendarId: event.CalendarId, Timed: event.Timed, Description: event.Description, Color: event.Color, CreatedAt: event.CreatedAt})
	}

	return events, err
}

func (e *Event) DeleteEvent(db *sql.DB) (id int, err error) {
	cmd := `delete from events where id = $1 RETURNING id`
	err = db.QueryRow(cmd, e.Id).Scan(&id)
	if err != nil {
		log.Fatalln(err)
	}

	return id, err
}
