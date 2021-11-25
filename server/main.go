package main

import (
	"log"
	"net"

	"github.com/shimo0108/task_list/server/services"

	pb "github.com/shimo0108/task_list/server/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Start listening port
	lis, err := net.Listen("tcp", ":9999")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register UsersServer to gRPC Server
	s := grpc.NewServer()

	pb.RegisterTasksServer(s, &services.TaskServer{})
	pb.RegisterCalendarsServer(s, &services.CalendarServer{})

	log.Println("start")
	//	pb.RegisterCalendarsServer(s, &services.CalendarService{})

	// Add grpc.reflection.v1alpha.ServerReflection
	reflection.Register(s)

	// Start server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
