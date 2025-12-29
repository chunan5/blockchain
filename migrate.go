package main

import (
	"blockchain/initializers"
	"blockchain/models"
	"log"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Cannot Migrate database")
	} else {
		log.Println("Database migrated")
	}

}
