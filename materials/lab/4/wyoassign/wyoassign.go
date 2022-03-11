package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"

)

type Response struct{
	Assignments []Assignment `json:"assignments"`
	Courses []Course `json:courses`
}

type Assignment struct {
	Id string `json:"id"`
	Course string `json:course`
	Title string `json:"title`
	Description string `json:"desc"`
	Points int `json:"points"`
	Due string `json:"due"`
}

type Course struct {
	Id string `json:"id"`
	Name string `json:"name`
	Description string `json:"desc"`
	Code int `json:"code"`
	Instructor string `json:"instr"`
}

var Assignments []Assignment
var Courses []Course
const Valkey string = "FooKey"

func InitAssignments(){
	var assignmnet Assignment
	assignmnet.Id = "Mike1A"
	assignmnet.Title = "Lab 4 "
	assignmnet.Description = "Some lab this guy made yesteday?"
	assignmnet.Points = 20
	assignmnet.Due = "Mar 11, 11:59pm"
	assignmnet.Course = "COSC 4010"
	Assignments = append(Assignments, assignmnet)
}

func InitCourses(){
	var course Course
	course.Id = "COSC 4010"
	course.Name = "CyberSecurity"
	course.Description = "Learn to hack & not get hacked"
	course.Code = 4010
	course.Instructor = "Mike"
	Courses = append(Courses, course)
}


func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func Home (w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to my awesome assignment site")
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Courses = Courses

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response.Courses)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}


func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	jsonResponse, err := json.Marshal(response.Assignments)

	if err != nil {
		return
	}

	//TODO 
	w.Write(jsonResponse)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)
	response["status"] = "No Such ID to GET"
	for _, assignment := range Assignments {
		if assignment.Id == params["id"]{
			json.NewEncoder(w).Encode(assignment)
			response["status"] = "Success"
			break
		}
	}
	//TODO : Provide a response if there is no such assignment
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func DeleteAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
			if assignment.Id == params["id"]{
				Assignments = append(Assignments[:index], Assignments[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func UpdateAssignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s Update end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// params := mux.Vars(r)
	// // var response Response
	// response := make(map[string]string)

	// response.Assignments = Assignments
	response := make(map[string]string)

	// var assignmnet Assignment
	r.ParseForm()

	response["status"] = "No Such ID to Update"
	for index, assign := range Assignments {
		if(assign.Id == r.FormValue("id")){
			assign.Title = r.FormValue("title")
			assign.Description =  r.FormValue("desc")
			assign.Points, _ =  strconv.Atoi(r.FormValue("points"))
			assign.Course = r.FormValue("course")
			assign.Due = r.FormValue("due")
			
			Assignments[index] = assign
			log.Println("assign: ", assign)
			log.Printf("Assignments: ", Assignments)
			w.WriteHeader(http.StatusAccepted)
			response["status"] = "Success"
			break
		}
	}
	w.WriteHeader(http.StatusNotFound)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}

func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignmnet Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging

	response := make(map[string]string)

	if(r.FormValue("id") != ""){
		assignmnet.Id =  r.FormValue("id")
		assignmnet.Title =  r.FormValue("title")
		assignmnet.Description =  r.FormValue("desc")
		assignmnet.Points, _ =  strconv.Atoi(r.FormValue("points"))
		assignmnet.Course = r.FormValue("course")
		assignmnet.Due = r.FormValue("due")
		Assignments = append(Assignments, assignmnet)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

	response["status"] = "Success"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var course Course
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging

	response := make(map[string]string)

	if(r.FormValue("id") != ""){
		course.Id =  r.FormValue("id")
		course.Name =  r.FormValue("name")
		course.Description =  r.FormValue("desc")
		course.Code, _ =  strconv.Atoi(r.FormValue("code"))
		course.Instructor = r.FormValue("instr")
		// course. = r.FormValue("due")
		Courses = append(Courses, course)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

	response["status"] = "Success"
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)

}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	
	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, course := range Courses {
			if course.Id == params["id"]{
				Courses = append(Courses[:index], Courses[index+1:]...)
				response["status"] = "Success"
				break
			}
	}
		
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

