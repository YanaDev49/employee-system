package database

import (
"fmt"
"log"
//"system/models"
"gorm.io/driver/postgres"
"gorm.io/gorm"
)


var Db *gorm.DB

func Connect() {

	ConnStr := "user=postgres dbname=employee-management sslmode=disable password=Password"

   var err error

	Db , err = gorm.Open(postgres.Open(ConnStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
		return
	}

	fmt.Println("Connection to database successful!")

    err = Db.AutoMigrate(&models.Department{}, &models.Employee{})
	if err != nil {
		log.Fatal("Error creating tables")
		return
	}
	

}