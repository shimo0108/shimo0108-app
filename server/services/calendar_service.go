package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/shimo0108/task_list/server/models"
	pb "github.com/shimo0108/task_list/server/proto"
)

type CalendarServer struct {
	pb.UnimplementedCalendarsServer
}

func (s *CalendarServer) FindAllCalendars(ctx context.Context, req *pb.FindAllCalendarsRequest) (*pb.FindAllCalendarsResponse, error) {
	cmd := `select id, name, visibility, created_at from calendars`
	rows, err := models.Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	calendar_lists := []*pb.Calendar{}
	for rows.Next() {
		var calendar models.Calendar
		err := rows.Scan(&calendar.Id, &calendar.Name, &calendar.Visibility, &calendar.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		var c_bool string
		var calendar_id string
		c_bool = strconv.FormatBool(calendar.Visibility)
		calendar_id = strconv.Itoa(calendar.Id)
		c := pb.Calendar{
			Id:         calendar_id,
			Name:       calendar.Name,
			Visibility: c_bool,
			CreatedAt:  calendar.CreatedAt.String(),
		}
		fmt.Println(calendar.Id, calendar.Name, calendar.Visibility, calendar.CreatedAt)

		calendar_lists = append(calendar_lists, &c)
	}

	lists := pb.FindAllCalendarsResponse{Calendar: calendar_lists}
	return &lists, nil
}

func (s *CalendarServer) FindByCalendarID(ctx context.Context, req *pb.FindByCalendarIDRequest) (*pb.FindByCalendarIDResponse, error) {
	var calendar_id int
	var c_bool string
	calendar_id, _ = strconv.Atoi(req.Id)
	c, _ := models.GetCalendar(calendar_id)
	time := c.CreatedAt.String()
	c_bool = strconv.FormatBool(c.Visibility)

	return &pb.FindByCalendarIDResponse{
		Calendar: &pb.Calendar{
			Id:         req.Id,
			Name:       c.Name,
			Visibility: c_bool,
			CreatedAt:  time,
		},
	}, nil
}

func (s *CalendarServer) CreateCalendar(ctx context.Context, req *pb.CreateCalendarRequest) (*pb.CreateCalendarResponse, error) {
	var calendar_id string

	c := &models.Calendar{}

	c_bool, _ := strconv.ParseBool(req.Calendar.Visibility)

	c.Name = req.Calendar.Name
	c.Visibility = c_bool
	c.CreateCalendar()

	calendar_id = strconv.Itoa(c.Id)

	return &pb.CreateCalendarResponse{
		Calendar: &pb.Calendar{
			Id:         calendar_id,
			Name:       c.Name,
			Visibility: req.Calendar.Visibility,
			CreatedAt:  c.CreatedAt.String(),
		},
	}, nil
}

func (s *CalendarServer) UpdateCalendar(ctx context.Context, req *pb.UpdateCalendarRequest) (*pb.UqdateCalendarResponse, error) {
	var calendar_id int
	calendar_id, _ = strconv.Atoi(req.Calendar.Id)

	c, _ := models.GetCalendar(calendar_id)

	if req.Calendar.Visibility == "" {
		req.Calendar.Visibility = strconv.FormatBool(c.Visibility)
	}
	c_bool, _ := strconv.ParseBool(req.Calendar.Visibility)

	c.Name = req.Calendar.Name
	c.Visibility = c_bool
	c.UpdateCalendar()

	return &pb.UqdateCalendarResponse{
		Calendar: &pb.Calendar{
			Id:         req.Calendar.Id,
			Name:       c.Name,
			Visibility: req.Calendar.Visibility,
			CreatedAt:  c.CreatedAt.String(),
		},
	}, nil
}

func (s *CalendarServer) DeleteCalendar(ctx context.Context, req *pb.DeleteCalendarRequest) (*pb.DeleteCalendarResponse, error) {
	var calendar_id int
	calendar_id, _ = strconv.Atoi(req.Id)
	c, _ := models.GetCalendar(calendar_id)
	c.DeleteCalendar()

	return &pb.DeleteCalendarResponse{
		Id: req.Id,
	}, nil
}
