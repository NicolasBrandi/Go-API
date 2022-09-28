package database

import (
	"log"
	"os"

	"github.com/NicolasBrandi/Go-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct{
	Db *gorm.DB
}

var Database DbInstance

//needs to have this exact name to work (ConnectDb, if connectDb does not work)
func ConnectDb(){ 
	//Innit the db
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	log.Println("Conection to the database succesful")
		db.Logger = logger.Default.LogMode(logger.Info)

	//Migrate the models to the db
	log.Println("Running migrations")
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})
		
	//set the db
	Database = DbInstance{Db:db}

}