package migrations

import (
	"log"
	"telmed_backend/db"
	"telmed_backend/models"
)

func Migrate() {
	database := db.GetDB()
	err := database.AutoMigrate(models.User{}, models.Doctor{}, models.Appointment{}, models.Prescription{})
	if err != nil {
		log.Println("Migrations failed", err)
		return
	}
	log.Println("Migrations ran successfully")

}
