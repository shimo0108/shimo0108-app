package services

import (
	"context"
	"fmt"
	"log"
	"strconv"

	pb "github.com/shimo0108/task_list/server/proto"

	"github.com/shimo0108/task_list/server/models"
)

type TaskServer struct {
	pb.UnimplementedTasksServer
}

func (s *TaskServer) FindAllTasks(ctx context.Context, req *pb.FindAllTasksRequest) (*pb.FindAllTasksResponse, error) {
	cmd := `select id, name, start_time, end_time, calendar_id, timed, COALESCE(description,''), COALESCE(color,''), created_at from tasks`
	rows, err := models.Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	task_lists := []*pb.Task{}
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Name, &task.StartTime, &task.EndTime, &task.CalendarId, &task.Timed, &task.Description, &task.Color, &task.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		var t_bool string
		var task_id string
		var calendar_id string
		t_bool = strconv.FormatBool(task.Timed)
		task_id = strconv.Itoa(task.Id)
		calendar_id = strconv.Itoa(task.CalendarId)
		t := pb.Task{
			Id:          task_id,
			Name:        task.Name,
			StartTime:   task.StartTime.String(),
			EndTime:     task.EndTime.String(),
			CalendarId:  calendar_id,
			Timed:       t_bool,
			Description: task.Description,
			Color:       task.Color,
			CreatedAt:   task.CreatedAt.String(),
		}
		fmt.Println(task.Id, task.Name, task.StartTime, task.EndTime, task.CalendarId, task.Timed, task.Description, task.Color, task.CreatedAt)

		task_lists = append(task_lists, &t)
	}

	lists := pb.FindAllTasksResponse{Task: task_lists}
	return &lists, nil
}

func (s *TaskServer) FindByTaskID(ctx context.Context, req *pb.FindByTaskIDRequest) (*pb.FindByTaskIDResponse, error) {
	var task_id int
	var t_bool string
	var calendar_id string
	task_id, _ = strconv.Atoi(req.Id)
	t, _ := models.GetTask(task_id)
	calendar_id = strconv.Itoa(t.CalendarId)
	time := t.CreatedAt.String()
	st_time := t.StartTime.String()
	e_time := t.EndTime.String()
	t_bool = strconv.FormatBool(t.Timed)

	return &pb.FindByTaskIDResponse{
		Task: &pb.Task{
			Id:          req.Id,
			Name:        t.Name,
			StartTime:   st_time,
			EndTime:     e_time,
			Timed:       t_bool,
			CalendarId:  calendar_id,
			Description: t.Description,
			Color:       t.Color,
			CreatedAt:   time,
		},
	}, nil
}

func (s *TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	var task_id string
	var calendar_id int
	calendar_id, _ = strconv.Atoi(req.Task.CalendarId)

	t := &models.Task{}
	st_time := stringToTime(req.Task.StartTime)
	end_time := stringToTime(req.Task.EndTime)

	t_bool, _ := strconv.ParseBool(req.Task.Timed)

	t.Name = req.Task.Name
	t.StartTime = st_time
	t.EndTime = end_time
	t.CalendarId = calendar_id
	t.Timed = t_bool
	t.Description = req.Task.Description
	t.Color = req.Task.Color
	t.CreateTask()

	task_id = strconv.Itoa(t.Id)

	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Id:          task_id,
			Name:        t.Name,
			StartTime:   req.Task.StartTime,
			EndTime:     req.Task.EndTime,
			CalendarId:  req.Task.CalendarId,
			Timed:       req.Task.Timed,
			Description: t.Description,
			Color:       t.Color,
			CreatedAt:   t.CreatedAt.String(),
		},
	}, nil
}

func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UqdateTaskResponse, error) {
	var task_id int
	var calendar_id int
	calendar_id, _ = strconv.Atoi(req.Task.CalendarId)

	task_id, _ = strconv.Atoi(req.Task.Id)
	t, _ := models.GetTask(task_id)

	if req.Task.Timed == "" {
		req.Task.Timed = strconv.FormatBool(t.Timed)
	}

	st_time := stringToTime(req.Task.StartTime)
	end_time := stringToTime(req.Task.EndTime)
	t_bool, _ := strconv.ParseBool(req.Task.Timed)

	t.Name = req.Task.Name
	t.StartTime = st_time
	t.EndTime = end_time
	t.CalendarId = calendar_id
	t.Timed = t_bool
	t.Description = req.Task.Description
	t.Color = req.Task.Color

	t.UpdateTask()

	return &pb.UqdateTaskResponse{
		Task: &pb.Task{
			Id:          req.Task.Id,
			Name:        t.Name,
			StartTime:   req.Task.StartTime,
			EndTime:     req.Task.EndTime,
			CalendarId:  req.Task.CalendarId,
			Timed:       req.Task.Timed,
			Description: t.Description,
			Color:       t.Color,
			CreatedAt:   t.CreatedAt.String(),
		},
	}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	var task_id int
	task_id, _ = strconv.Atoi(req.Id)
	t, _ := models.GetTask(task_id)
	t.DeleteTask()

	return &pb.DeleteTaskResponse{
		Id: req.Id,
	}, nil
}
