package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Model for courses - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middlewares, helper -file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("server in go")

	//get all courses
	//create new courses
	//delete the courses
	//update the courses
	r := mux.NewRouter()


	//dummy data
	courses = append(courses, Course{CourseId: "2", CourseName: "react", CoursePrice: 299, Author: &Author{Fullname: "Hruday Chowdary", Website: "hruday.is-a-dev"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "DSA", CoursePrice: 199, Author: &Author{Fullname: "Hasini Chowdary", Website: "coursera"}})

	//routes
	r.HandleFunc("/",serveHome).Methods("GET") //working
	r.HandleFunc("/courses",getAllCourses).Methods("GET") //working
	r.HandleFunc("/course/{courseid}",getOneCourse).Methods("GET") //working
	r.HandleFunc("/course",createOneCourse).Methods("POST") //working
	r.HandleFunc("/course/{courseid}",updateOneCourse).Methods("PUT") //working
	r.HandleFunc("/course/{courseid}",deleteOneCourse).Methods("DELETE") //working
	//boom babyyy



	log.Fatal(http.ListenAndServe(":4000",r))
}

//controllers - file

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api of Hruday</h1>"))
}

// get all courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the courses and find the matching id from the params and the database
	// and return the response

	for _, course := range courses {
		if course.CourseId == params["courseid"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with response id")
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what is req body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
	}

	//what about data is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside the req body")
		return
	}

	//generate unique id of type string
	course.CourseId = strconv.Itoa(rand.Intn(100))
	//append course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	
	
	//loop through find and update(remove and update)
	for index, course := range courses {
		if course.CourseId == params["courseid"] {
			courses = append(courses[:index],courses[index+1:]...) //remove

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["courseid"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if(course.CourseId == params["courseid"]) {
			courses = append(courses[:index],courses[index+1:]...)

			json.NewEncoder(w).Encode("Course deleted successfully")
			break
		}
	}
}
