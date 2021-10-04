package main

import (
    "context"
    "log"
    "math/rand"
    "net"

    pb "example.com/go-apiservice-grpc/apiservice"
    "google.golang.org/grpc"
)

const (
    port = ":50051"
)

type ApiServiceServer struct {
    pb.UnimplementedRouteCourseServer
}

func (s *ApiServiceServer) AddCourse(ctx context.Context, in *pb.NewCourse) (*pb.Course, error) {
    log.Printf("Received: %v", in.GetName())
    var courseId int64 = int64(rand.Intn(1000))
    return &pb.Course{Name: in.GetName(), Id: courseId}, nil
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    server := grpc.NewServer()
    pb.RegisterRouteCourseServer(server, &ApiServiceServer{})
    log.Printf("server listening at %v", lis.Addr())
    if err := server.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}