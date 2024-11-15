package main

import (
	"log"
	"net/http"

	"github.com/Stuko0/PendejadaUni/internal/database"
	"github.com/Stuko0/PendejadaUni/pkg/models"
	"github.com/Stuko0/PendejadaUni/pkg/routes"
)

func main() {
	database.Connect()
	database.DB.AutoMigrate(&models.Student{})
	database.DB.AutoMigrate(&models.Classroom{})
	r := routes.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}