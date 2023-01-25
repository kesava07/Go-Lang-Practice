package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// models -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB

var courses []Course

// Middleware, Helpers

func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""

}

func main() {
	fmt.Println("API running on PORT:4000")
	r := mux.NewRouter()

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", serveCourses).Methods("GET")
	r.HandleFunc("/course/{id}", serveCourseById).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourseById).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourseById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to API by Golang </h1>"))
}

func serveCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func serveCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get course by Id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found with the given Id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// if BODY is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Information is not enough")
		return
	}

	// If course is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Course name is required")
		return
	}

	// Generate unique id
	// Append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode("Course created successfully")
	return
}

func updateCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course by ID")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	if r.Body == nil {
		json.NewEncoder(w).Encode("Course data is required")
	}

	if params["id"] == "" {
		json.NewEncoder(w).Encode("Course Id required to update a course")
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("Updated course successfully")
			return
		}
	}
}

func deleteCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update course by ID")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	if params["id"] == "" {
		json.NewEncoder(w).Encode("Bad request")
	}

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Deleted course successfully")
			break
		}
	}

}
