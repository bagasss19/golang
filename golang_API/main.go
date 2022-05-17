package main

import (
	"golang_api/config"
	"golang_api/controller"
	"golang_api/middleware"
	repository "golang_api/repository/postgres"
	"golang_api/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	err    error
)

func readEnvironmentFile() {
	//Environment file Load --------------------------------
	err := godotenv.Load(".secret.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(3)
	}
}

func main() {
	readEnvironmentFile()
	DBConn, err = config.DBConnect()
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}

	// Init Repository
	userRepo := repository.NewUserRepository(DBConn)

	// Init Usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Init Controller
	userController := controller.NewUserController(userUsecase)

	r := gin.Default()
	r.Use(middleware.CORS())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := r.Group("/api")

	// User Route
	user := api.Group("/user")
	user.GET("/crud", middleware.TokenAuthMiddleware(), userController.GetAllUser)
	user.POST("/crud", userController.CreateUser)
	user.GET("/crud/:id", userController.GetUserByID)
	user.DELETE("/crud/:id", userController.DeleteUser)
	user.PUT("/crud/:id", userController.UpdateUser)
	user.POST("/login", userController.Login)

	r.Run() // listen and serve on 0.0.0.0:8080
}
