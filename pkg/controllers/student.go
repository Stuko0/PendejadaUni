package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Stuko0/PendejadaUni/internal/database"
	"github.com/Stuko0/PendejadaUni/pkg/models"
	"github.com/gorilla/mux"
)

func GetStudents(w http.ResponseWriter, r *http.Request){
	var students []models.Student
	database.DB.Find(&students)
	json.NewEncoder(w).Encode(students)
}

func CreateStudent(w http.ResponseWriter, r *http.Request){
	var student models.Student
	json.NewDecoder(r.Body).Decode(&student)
	database.DB.Create(&student)
	json.NewEncoder(w).Encode(student)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var student models.Student
	database.DB.First(&student,params["id"])
	json.NewDecoder(r.Body).Decode(&student)
	database.DB.Save(&student)
	json.NewEncoder(w).Encode(student)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var student models.Student
	database.DB.Delete(&student,params["id"])
	json.NewEncoder(w).Encode("Student Deleted")
}