package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"wyoassign/wyoassign"
)


func main() {
	wyoassign.InitAssignments()
	wyoassign.InitCourses()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints

	router.HandleFunc("/", wyoassign.Home).Methods("GET")
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.GetAssignment).Methods("GET")
	router.HandleFunc("/delete-assignment/{id}", wyoassign.DeleteAssignment).Methods("DELETE")		
	router.HandleFunc("/create-assignment", wyoassign.CreateAssignment).Methods("POST")	
	router.HandleFunc("/update-assignment", wyoassign.UpdateAssignment).Methods("POST")
	router.HandleFunc("/courses", wyoassign.GetCourses).Methods("GET")
	router.HandleFunc("/create-course", wyoassign.CreateCourse).Methods("POST")	
	router.HandleFunc("/delete-course/{id}", wyoassign.DeleteCourse).Methods("DELETE")		

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}