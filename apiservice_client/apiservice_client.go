package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-apiservice-grpc/apiservice"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRouteCourseClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newCourses = make(map[string]string)
	newCourses["Test"] = "TestCourse"
	for name := range newCourses {
		res, err := client.AddCourse(ctx, &pb.NewCourse{ Name: name })
		if err != nil {
			log.Fatalf("could not create course: %v", err)
		}
		log.Printf(`Course Details:
NAME: %s
ID: %d`, res.GetName(), res.GetId())
	}
}