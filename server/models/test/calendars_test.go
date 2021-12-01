package test

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shimo0108/shimo0108-app/server/handler"
	"github.com/shimo0108/shimo0108-app/server/models"
	"github.com/stretchr/testify/assert"
)

func TestGetCalendar(t *testing.T) {
	// モックDBの初期化
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to init db mock")
	}
	defer db.Close()

	id := 1

	// dbドライバに対する操作のモック定義
	columns := []string{"id", "name", "visibility", "color", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta("select id, name, visibility, COALESCE(color,''), created_at from calendars where id = $1")). // expectedSQL: 想定される実行クエリをregexpで指定（指定文字列が含まれるかどうかを見る）
																	WithArgs(id).                                                                                                      // 想定されるプリペアドステートメントへの引数
																	WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "勉強", true, "red", handler.StringToTime("2006-01-02 15:04:05"))) // 返戻する行情報の指定

	calendar, err := models.GetCalendar(id, db)
	if err != nil {
		t.Fatalf("failed to get calendar: %s", err)
	}
	fmt.Printf("%v", calendar)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("failed to ExpectationWerMet(): %s", err)
	}

	assert.Equal(t, err, nil)
	assert.Equal(t, calendar.Id, id)
	assert.Equal(t, calendar.Name, "勉強")
	assert.Equal(t, calendar.Visibility, true)
	assert.Equal(t, calendar.Color, "red")
}

func TestCreateCalendar(t *testing.T) {
	t.Run(
		"Create",
		func(t *testing.T) {
			// Arrange
			ca := &models.Calendar{}
			var id int = 1

			name := "予定"
			visibility := true
			color := "red"

			ca.Name = name
			ca.Visibility = visibility
			ca.Color = color
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			columns := []string{"id"}

			mock.ExpectQuery(regexp.QuoteMeta(`insert into calendars (name, visibility, color, created_at) values ($1, $2, $3, $4) RETURNING id`)).
				WithArgs(name, visibility, color, AnyTime{}).
				WillReturnRows(sqlmock.NewRows(columns).AddRow(id))

			res, err := ca.CreateCalendar(db)
			if err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, err, nil)
			assert.Equal(t, res, id)
			assert.Equal(t, ca.Name, "予定")
			assert.Equal(t, ca.Visibility, visibility)
			assert.Equal(t, ca.Color, "red")

		},
	)
}

func TestUpdateCalendar(t *testing.T) {
	t.Run(
		"Update",
		func(t *testing.T) {
			// Arrange
			ca := &models.Calendar{}

			id := 1
			name := "予定"
			visibility := true
			color := "red"

			ca.Id = id
			ca.Name = name
			ca.Visibility = visibility
			ca.Color = color
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta(`update calendars
						set name = $1
						,visibility = $2
						,color = $3
						where id = $4`)).
				WithArgs(name, visibility, color, id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			ca.UpdateCalendar(db)

			assert.Equal(t, err, nil)
			assert.Equal(t, ca.Id, 1)
			assert.Equal(t, ca.Name, "予定")
			assert.Equal(t, ca.Visibility, visibility)
			assert.Equal(t, ca.Color, "red")

			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}

func TestDeleteCalendar(t *testing.T) {
	t.Run(
		"Delete",
		func(t *testing.T) {
			id := 1
			ca := &models.Calendar{}
			ca.Id = id
			e := &models.Event{}
			e.CalendarId = id

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Error(err.Error())
			}
			defer db.Close()
			mock.ExpectExec(regexp.QuoteMeta(`delete from events where calendar_id = $1`)).
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectExec(regexp.QuoteMeta(`delete from calendars where id = $1`)).
				WithArgs(id).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err = ca.DeleteCalendar(db)

			assert.Equal(t, err, nil)

			if err != nil {
				t.Error(err.Error())
			}
		},
	)
}
