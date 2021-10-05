package swagger

import "fmt"

var database Database

type Database struct {
	courses  []Course
	students []Student
	teachers []Teacher
}

func GetDB() *Database {
	return &database
}

func CreateDB() {
	database = Database{make([]Course, 0), make([]Student, 0), make([]Teacher, 0)}
}

func DBAddCourse(c Course) bool {
	if (CourseIdExists(c.Id)) {
		return false
	}
	database.courses = append(database.courses, c)
	return true
}

func DBDeleteCourse(cid int64) bool {
	index := IndexFromId(cid)
	if (index == -1) {
		return false
	}
	database.courses[index] = database.courses[len(database.courses)-1]
	database.courses = database.courses[:len(database.courses)-1]
	return true
}

func DBGetCourse(cid int64) (Course, error) {
	index := IndexFromId(cid)
	if (index == -1) {
		return Course{}, fmt.Errorf("Course not found")
	}
	return database.courses[index], nil
}

func DBUpdateCourse(cid int64, c Course) error {
	index := IndexFromId(cid)
	if (index == -1) {
		return fmt.Errorf("Course not found")
	}
	if (c.Id != cid && CourseIdExists(c.Id)) {
		return fmt.Errorf("ID Already exists")
	}
	database.courses[index] = c
	return nil
}

func DBGetCourseOverview() []Course {
	return database.courses
}


func CourseIdExists(id int64) bool {
	for _, course := range database.courses {
        if (course.Id == id) {
			return true
		}
    }
	return false
}

func IndexFromId(id int64) int {
	for idx, course := range database.courses {
        if (course.Id == id) {
			return idx
		}
    }
	return -1
}