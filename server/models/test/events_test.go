package test

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shimo0108/shimo0108-app/server/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	t.Run(
		"Create",
		func(t *testing.T) {
			// Arrange
			e := &models.Event{}

			var id int = 1
			name := "Goの勉強"
			color := "red"
			started_at := stringToTime("2006-01-02 15:04:05")
			ended_at := stringToTime("2006-01-02 15:04:05")
			calendar_id := 3
			timed := true
			description := "テストを書く"

			e.Name = name
			e.StartedAt = started_at
			e.EndedAt = ended_at
			e.CalendarId = calendar_id
			e.Timed = timed
			e.Description = description
			e.Color = color
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			columns := []string{"id"}

			mock.ExpectQuery(regexp.QuoteMeta(`insert into events (name, started_at, ended_at, calendar_id, timed, description, color, created_at) values ($1, $2, $3, $4, $5, $6 ,$7, $8) RETURNING id`)).
				WithArgs(name, started_at, ended_at, calendar_id, timed, description, color, AnyTime{}).
				WillReturnRows(sqlmock.NewRows(columns).AddRow(id))

			res, err := e.CreateEvent(db)
			if err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, err, nil)
			assert.Equal(t, res, id)
			assert.Equal(t, e.Name, name)
			assert.Equal(t, e.StartedAt, started_at)
			assert.Equal(t, e.EndedAt, ended_at)
			assert.Equal(t, e.CalendarId, calendar_id)
			assert.Equal(t, e.Timed, timed)
			assert.Equal(t, e.Description, description)
			assert.Equal(t, e.Color, color)

		},
	)
}

func TestUpdateEvent(t *testing.T) {
	t.Run(
		"Update",
		func(t *testing.T) {
			// Arrange
			e := &models.Event{}

			id := 1
			name := "Goの勉強"
			color := "red"
			started_at := stringToTime("2006-01-02 15:04:05")
			ended_at := stringToTime("2006-01-02 15:04:05")
			calendar_id := 3
			timed := true
			description := "テストを書く"

			e.Id = id
			e.Name = name
			e.StartedAt = started_at
			e.EndedAt = ended_at
			e.CalendarId = calendar_id
			e.Timed = timed
			e.Description = description
			e.Color = color
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta(`update events set name = $1 ,started_at = $2 ,ended_at = $3 ,calendar_id = $4 ,timed = $5 ,description = $6 ,color = $7 where id = $8`)).
				WithArgs(name, started_at, ended_at, calendar_id, timed, description, color, id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			e.UpdateEvent(db)

			assert.Equal(t, err, nil)
			assert.Equal(t, e.Id, id)
			assert.Equal(t, e.Name, name)
			assert.Equal(t, e.StartedAt, started_at)
			assert.Equal(t, e.EndedAt, ended_at)
			assert.Equal(t, e.CalendarId, calendar_id)
			assert.Equal(t, e.Timed, timed)
			assert.Equal(t, e.Description, description)
			assert.Equal(t, e.Color, color)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}

func TestDeleteEvent(t *testing.T) {
	t.Run(
		"Delete",
		func(t *testing.T) {
			// Arrange
			e := &models.Event{}
			id := 1
			e.Id = id

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			columns := []string{"id"}
			mock.ExpectQuery(regexp.QuoteMeta(`delete from events where id = $1`)).
				WithArgs(id).
				WillReturnRows(sqlmock.NewRows(columns).AddRow(id))

			res, err := e.DeleteEvent(db)
			if err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, err, nil)
			assert.Equal(t, res, id)

		},
	)
}
