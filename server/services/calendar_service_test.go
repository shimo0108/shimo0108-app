package services

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	pb "github.com/shimo0108/task_list/server/proto"
)

func TestFindByCalendarID(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCalendarsClient(conn)
	resp, err := client.FindByCalendarID(ctx, &pb.FindByCalendarIDRequest{Id: "1"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Calendar.Id != "1" {
		t.Fatal("The Calendar ID must be 1.")
	}
}

func TestCreateCalendar(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCalendarsClient(conn)

	calendar := &pb.Calendar{
		Id:         "",
		Name:       "行事",
		Visibility: "true",
		CreatedAt:  "",
	}

	resp, err := client.CreateCalendar(ctx, &pb.CreateCalendarRequest{Calendar: calendar})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetCalendar().Name != "行事" {
		t.Fatal("The reply must be '行事'")
	}
	if resp.GetCalendar().Visibility != "true" {
		t.Fatal("Wrong Visibility value")
	}
}

func TestUpdateCalendar(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCalendarsClient(conn)

	calendar := &pb.Calendar{
		Id:         "1",
		Name:       "予定",
		Visibility: "true",
		CreatedAt:  "",
	}

	resp, err := client.UpdateCalendar(ctx, &pb.UpdateCalendarRequest{Calendar: calendar})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetCalendar().Name != "予定" {
		t.Fatal("The reply must be '行事'")
	}
	if resp.GetCalendar().Visibility != "true" {
		t.Fatal("Wrong Visibility value")
	}
}

func TestDeleteCalendar(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewCalendarsClient(conn)
	resp, err := client.DeleteCalendar(ctx, &pb.DeleteCalendarRequest{Id: "2"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.GetId() != "2" {
		t.Fatal("The Calendar ID must be 2.")
	}
}
