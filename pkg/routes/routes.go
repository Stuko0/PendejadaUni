package routes

import (
	"github.com/Stuko0/PendejadaUni/pkg/controllers"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/students", controllers.GetStudents).Methods("GET")
	r.HandleFunc("/students", controllers.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", controllers.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", controllers.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/classroom", controllers.GetClassrooms).Methods("GET")
	r.HandleFunc("/classroom", controllers.CreateClassroom).Methods("POST")
	r.HandleFunc("/classroom/{id}", controllers.UpdateClassroom).Methods("PUT")
	r.HandleFunc("/classroom/{id}", controllers.DeleteClassroom).Methods("DELETE")
	return r
}
