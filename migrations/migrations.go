package migrations

import (
	"log"
	"telmed_backend/db"
)

func Migrate() {
	database := db.GetDB()
	err := database.AutoMigrate()
	if(err!=nil) {
		log.Println("Migrations failed",err)
		return 
	}
	log.Println("Migrations ran successfully")
	

}
