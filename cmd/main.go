package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafmasloman/snappeeps_backend/internal/config"
	"github.com/rafmasloman/snappeeps_backend/internal/domain"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	fmt.Println("Started golang apps")

	db, _ := config.InitializeDatabase()

	db.AutoMigrate(&domain.User{})

	router.Run(":8091")
}
