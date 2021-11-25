package services

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "github.com/shimo0108/task_list/server/proto"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterTasksServer(s, &TaskServer{})
	pb.RegisterCalendarsServer(s, &CalendarServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func bufDialer(ctx context.Context, address string) (net.Conn, error) {
	return lis.Dial()
}

func TestFindByTaskID(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTasksClient(conn)
	resp, err := client.FindByTaskID(ctx, &pb.FindByTaskIDRequest{Id: "1"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Task.Id != "1" {
		t.Fatal("The task ID must be 1.")
	}
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTasksClient(conn)

	task := &pb.Task{
		Id:          "",
		Name:        "Goの勉強",
		StartTime:   "2006-01-02 15:04:05",
		EndTime:     "2006-01-03 15:04:05",
		CalendarId:  "1",
		Timed:       "true",
		Description: "早めに終わらせる",
		Color:       "red",
		CreatedAt:   "",
	}

	resp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{Task: task})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetTask().Name != "Goの勉強" {
		t.Fatal("The reply must be 'Goの勉強'")
	}
	if resp.GetTask().StartTime != "2006-01-02 15:04:05" {
		t.Fatal("Wrong StartTime value")
	}
	if resp.GetTask().EndTime != "2006-01-03 15:04:05" {
		t.Fatal("Wrong EndTime value")
	}
	if resp.GetTask().CalendarId != "1" {
		t.Fatal("The reply calendar_id must be '1'")
	}
	if resp.GetTask().Timed != "true" {
		t.Fatal("Wrong Timed value")
	}
	if resp.GetTask().Description != "早めに終わらせる" {
		t.Fatal("The reply must be '早めに終わらせる'")
	}
	if resp.GetTask().Color != "red" {
		t.Fatal("The reply must be 'red'")
	}
}

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTasksClient(conn)

	task := &pb.Task{
		Id:          "1",
		Name:        "gRPCの勉強",
		StartTime:   "2006-01-03 15:04:05",
		EndTime:     "2006-01-05 15:04:05",
		CalendarId:  "3",
		Timed:       "false",
		Description: "docker上で動作を確認",
		Color:       "red",
	}

	resp, err := client.UpdateTask(ctx, &pb.UpdateTaskRequest{Task: task})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetTask().Name != "gRPCの勉強" {
		t.Fatal("The reply must be 'gRPCの勉強'")
	}
	if resp.GetTask().StartTime != "2006-01-03 15:04:05" {
		t.Fatal("Wrong StartTime value")
	}
	if resp.GetTask().EndTime != "2006-01-05 15:04:05" {
		t.Fatal("Wrong EndTime value")
	}
	if resp.GetTask().CalendarId != "3" {
		t.Fatal("The reply calendar_id must be '3'")
	}
	if resp.GetTask().Timed != "false" {
		t.Fatal("Wrong Timed value")
	}
	if resp.GetTask().Description != "docker上で動作を確認" {
		t.Fatal("The reply must be 'docker上で動作を確認'")
	}
	if resp.GetTask().Color != "red" {
		t.Fatal("The reply must be 'red'")
	}
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewTasksClient(conn)
	resp, err := client.DeleteTask(ctx, &pb.DeleteTaskRequest{Id: "1"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetId() != "1" {
		t.Fatal("The task ID must be 1.")
	}
}
