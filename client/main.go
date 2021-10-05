package main

import (
	"example.com/restclient/client"
	"example.com/restclient/client/course"
	"example.com/restclient/models"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

func main() {
	transport := httptransport.New("localhost:8080", "", []string{"http"})
	c := client.New(transport, strfmt.Default)

	params := course.NewAddCourseParams()
	params.Body = &models.Course{ID: 3, Name: "testcourse"}

	courseres, err := c.Course.AddCourse(params)
	if err != nil {
		fmt.Printf("could not add course: %v", err)
		return 
	}
	fmt.Printf("New Course: %v", courseres.Error())
}
