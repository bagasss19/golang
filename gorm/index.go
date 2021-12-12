package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID              int
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	Username        string `json:"username" binding:"required"`
	DatetimeCreated time.Time
	DatetimeUpdated time.Time
}

func getUser(c *gin.Context) {
	dsn := "host=localhost user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	var user []User
	db.Find(&user)
	// fmt.Print(result.RowsAffected, "<<????")
	c.JSON(200, gin.H{
		"message": user,
	})
}

func Register(c *gin.Context) {
	dsn := "host=localhost user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// Validate input
	createuser := User{DatetimeCreated: time.Now(), DatetimeUpdated: time.Now()}

	if err := c.ShouldBindJSON(&createuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&createuser)
	c.JSON(201, gin.H{"status": createuser})
}

func updateUser(c *gin.Context) {
	dsn := "host=localhost user=postgres password=postgres dbname=golang port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("failed to connect database")
	}

	// Validate input
	createuser := User{DatetimeUpdated: time.Now()}
	id := c.Param("id")

	if err := c.ShouldBindJSON(&createuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Find()

	db.Create(&createuser)
	c.JSON(201, gin.H{"status": createuser})
}

func main() {
	r := gin.Default()
	r.GET("/user", getUser)
	r.POST("/user", Register)
	r.POST("/user/:id", updateUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
