package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Stuko0/PendejadaUni/internal/database"
	"github.com/Stuko0/PendejadaUni/pkg/models"
	"github.com/gorilla/mux"
)

func GetClassrooms(w http.ResponseWriter, r *http.Request){
	var classrooms []models.Classroom
	database.DB.Find(&classrooms)
	json.NewEncoder(w).Encode(classrooms)
}

func CreateClassroom(w http.ResponseWriter, r *http.Request){
	var classroom []models.Classroom
	json.NewDecoder(r.Body).Decode(&classroom)
	database.DB.Create(&classroom)
	json.NewEncoder(w).Encode(classroom)
}

func UpdateClassroom(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var classroom []models.Classroom
	database.DB.First(&classroom,params["id"])
	json.NewDecoder(r.Body).Decode(&classroom)
	database.DB.Save(&classroom)
	json.NewEncoder(w).Encode(classroom)
}

func DeleteClassroom(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	var classroom []models.Classroom
	database.DB.Delete(&classroom,params["id"])
	json.NewEncoder(w).Encode("Classroom Deleted")
}