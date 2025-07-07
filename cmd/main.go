package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafmasloman/snappeeps_backend/internal/config"
	"github.com/rafmasloman/snappeeps_backend/internal/delivery/http"
	"github.com/rafmasloman/snappeeps_backend/internal/domain"
	"github.com/rafmasloman/snappeeps_backend/internal/infrastructure/mysql"
	"github.com/rafmasloman/snappeeps_backend/internal/usecase"
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

	userRepository := mysql.NewUserRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepository)
	authHttp := http.NewAuthHttp(authUsecase)

	apiRouter := router.Group("/api/v1")
	{
		authRouter := apiRouter.Group("/auth")
		authRouter.POST("/register", authHttp.Register)
	}

	router.Run(":8091")
}
