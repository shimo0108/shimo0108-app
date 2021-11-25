package models

import (
	"log"
	"time"
)

type Task struct {
	Id          int
	Name        string
	StartTime   time.Time
	EndTime     time.Time
	CalendarId  int
	Timed       bool
	Description string
	Color       string
	CreatedAt   time.Time
}

func (t *Task) CreateTask() (err error) {
	cmd := `insert into tasks (name, start_time, end_time, calendar_id, timed, description, color, created_at) values ($1, $2, $3, $4, $5, $6 ,$7, $8)`

	_, err = Db.Exec(cmd,
		t.Name,
		t.StartTime,
		t.EndTime,
		t.CalendarId,
		t.Timed,
		t.Description,
		t.Color,
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTask(id int) (task Task, err error) {
	task = Task{}
	cmd := `select id, name, start_time, end_time, calendar_id, timed, description, color, created_at
	from tasks where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&task.Id,
		&task.Name,
		&task.StartTime,
		&task.EndTime,
		&task.CalendarId,
		&task.Timed,
		&task.Description,
		&task.Color,
		&task.CreatedAt,
	)
	return task, err
}

func (t *Task) UpdateTask() (err error) {
	cmd := `update tasks
	set name = $1
	,start_time = $2
	,end_time = $3
	,calendar_id = $4
	,timed = $5
	,description = $6
	,color = $7 where id = $8`
	_, err = Db.Exec(cmd, t.Name, t.StartTime, t.EndTime, t.CalendarId, t.Timed, t.Description, t.Color, t.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Task) DeleteTask() (err error) {
	cmd := `delete from tasks where id = $1`
	_, err = Db.Exec(cmd, t.Id)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
