/*
 * ITU API
 *
 * ITU REST API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ITU Student, Teacher and Course Management")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"AddCourse",
		strings.ToUpper("Post"),
		"/course",
		AddCourse,
	},

	Route{
		"DeleteCourse",
		strings.ToUpper("Delete"),
		"/course/{courseId}",
		DeleteCourse,
	},

	Route{
		"GetCourseById",
		strings.ToUpper("Get"),
		"/course/{courseId}",
		GetCourseById,
	},

	Route{
		"UpdateCourse",
		strings.ToUpper("Put"),
		"/course/{courseId}",
		UpdateCourse,
	},

	Route{
		"GetOverview",
		strings.ToUpper("Get"),
		"/courses",
		GetOverview,
	},

	Route{
		"AddStudent",
		strings.ToUpper("Post"),
		"/student",
		AddStudent,
	},

	Route{
		"DeleteStudent",
		strings.ToUpper("Delete"),
		"/student/{studentId}",
		DeleteStudent,
	},

	Route{
		"FindStudentByStatus",
		strings.ToUpper("Get"),
		"/student/findByStatus",
		FindStudentByStatus,
	},

	Route{
		"GetStudentById",
		strings.ToUpper("Get"),
		"/student/{studentId}",
		GetStudentById,
	},

	Route{
		"UpdateStudent",
		strings.ToUpper("Put"),
		"/student",
		UpdateStudent,
	},

	Route{
		"UploadStudentPhoto",
		strings.ToUpper("Post"),
		"/student/{studentId}/uploadImage",
		UploadStudentPhoto,
	},

	Route{
		"AddTeacher",
		strings.ToUpper("Post"),
		"/teacher",
		AddTeacher,
	},

	Route{
		"DeleteTeacher",
		strings.ToUpper("Delete"),
		"/teacher/{teacherId}",
		DeleteTeacher,
	},

	Route{
		"FindTeacherByStatus",
		strings.ToUpper("Get"),
		"/teacher/findByStatus",
		FindTeacherByStatus,
	},

	Route{
		"GetTeacherById",
		strings.ToUpper("Get"),
		"/teacher/{teacherId}",
		GetTeacherById,
	},

	Route{
		"UpdateTeacher",
		strings.ToUpper("Put"),
		"/teacher",
		UpdateTeacher,
	},

	Route{
		"UploadTeacherPhoto",
		strings.ToUpper("Post"),
		"/teacher/{teacherId}/uploadImage",
		UploadTeacherPhoto,
	},
}
