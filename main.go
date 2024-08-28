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

// creation of fakeDB
var courses []Course

//middleware or helper method that will basically do not allow someone to add course details if courseId and courseName is empty

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""

	//allowing user to move forward if the courseID is empty
	return c.CourseName == ""

}

func main() {
	fmt.Println("API- working properly...")
	r:=mux.NewRouter();
	
	//seeding 
	courses=append(courses, Course{CourseId: "2",CourseName: "React.js",CoursePrice: 399,Author: &Author{Fullname: "Suvam Kumar Nanda",Website: "learning.com"}})
	courses=append(courses, Course{CourseId: "4",CourseName: "Fullstack java",CoursePrice: 499,Author: &Author{Fullname: "Suvam Kumar Nanda",Website: "learning.com"}})
    
    //routing part
	r.HandleFunc("/",serveHome).Methods("GET");
	r.HandleFunc("/courses",getAllCourses).Methods("GET");  
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET");  //here we have to parameters id as we have taken params["id"] in getOneCourse controller
    r.HandleFunc("/create",createOneCourse).Methods("POST");
	r.HandleFunc("/update/{id}",updateOneCourse).Methods("PUT");
	r.HandleFunc("/delete/{id}",deleteOneCourse).Methods("DELETE");
	

	


	
	
	//listen on a port
	log.Fatal(http.ListenAndServe(":4000",r));

}

//controllers-file

// serveHome (code for  homepage)
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang suvam</h1>"))
}

// getAll course
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	//set header of the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get a single course based on the id
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course with it's id ")
	//set the header to the response
	w.Header().Set("Content-Type", "application/json")

	//grab id from request and here we are not only accessing id but all the variables so we have given the name as params

	params := mux.Vars(r)

	//loop through courses , find matching id and return the respose
	//use for-range loop

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
		json.NewEncoder(w).Encode("No course is found with the id")

	}
}

// create a course
func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if body is empty:
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about {}
	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return

	}
	// //appending the data to the slice
	// courses=append(courses,course);

	//generate the unique Id , convert the id to string (here we don't want to rely on user for courseID want to generate courseID automatically)
	//after that append new courses to the courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	//Itoa() is  a method in  strconv package that basically converts the integer to string
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)
	return
}



//update a course details
func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Update one course");
   w.Header().Set("Content-Type","application/json");

   //grab the id from request
   params:=mux.Vars(r);

   //loop through the slice ,get the id matched,remove the value ,add the value with myId (the id that we have grabed  in paraams)
   for index,course:=range courses{
    if course.CourseId==params["id"] {
		courses= append(courses[:index],courses[index+1:]... );
		var course Course;
	    _=json.NewDecoder(r.Body).Decode(&course);  //here we are decoding  the details of the request.Body and storing them in new course variable declared just above
		
		course.CourseId=params["id"];   //here we are setting the id of the new course with the id that we have grabbed from the Request
        courses=append(courses, course) ; //adding the updated course details to the original slice
         json.NewEncoder(w).Encode(course);
		 return;


	}
   }
   


   

   }

   //Delete One course

   func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("Delete one course");
	w.Header().Set("Content-Type","application/json");

	//garb the unique Id from the request 
	params:=mux.Vars(r);

	//loop through the slice,remove(index,index+1)
	for index,course:=range courses{
		if course.CourseId==params["id"] {
			courses=append(courses[:index],courses[index+1:]... );
			json.NewEncoder(w).Encode("Course is deleted");
			break;
			
		}
	}


   }